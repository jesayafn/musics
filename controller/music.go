package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jessie-txt/musics/configs"
	"go.mongodb.org/mongo-driver/bson"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

type MusicsDetail struct {
	Id             string   `json:"id"`
	Singer         []string `json:"singer"`
	SongName       string   `json:"songName"`
	Album          string   `json:"album"`
	Release        string   `json:"release"`
	RecordingLabel string   `json:"recordingLabel"`
}

type MusicOverview []struct {
	Id       string   `json:"id"`
	Singer   []string `json:"singer"`
	SongName string   `json:"songName"`
	Album    string   `json:"album"`
}

type Response struct {
	Success   string      `json:"success"`
	Message   string      `json:"message"`
	ErrorCode int         `json:"errorCode"`
	Data      interface{} `json:"data"`
}

var (
	timeNow       = time.Now().String()
	client, err   = configs.MongoDb()
	ignoreDeleted = bson.D{{Key: "$or",
		Value: bson.A{
			bson.D{{
				Key: "deleted",
				Value: bson.D{{
					Key:   "$exists",
					Value: false}},
			}},
			bson.D{{
				Key: "deleted",
				Value: bson.D{{
					Key:   "$exists",
					Value: true}, {
					Key: "$nin",
					Value: bson.A{
						true,
					},
				}},
			}},
		},
	}}
)

func GetMusics(c *gin.Context) {
	var music MusicOverview
	// var empty MusicOverview

	opts := mongoOptions.Find().SetProjection(bson.D{{Key: "deleted", Value: 0}})
	collection, err := configs.MongoDbCollection(client, "musics", "music").Find(context.TODO(),
		ignoreDeleted, opts)

	if err != nil {
		configs.MongoDbLogger()
		log.Fatalln(err)
	}
	err = collection.All(context.TODO(), &music)

	if err != nil {
		configs.MongoDbLogger()
		c.JSON(http.StatusInternalServerError, Response{
			Success:   "false",
			Message:   "Internal Server Error",
			ErrorCode: 2,
			Data:      "null",
		})
	} else if len(music) == 0 {
		c.JSON(http.StatusNotFound, Response{
			Success:   "false",
			Message:   "Not Found",
			ErrorCode: 1,
			Data:      "null",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Success:   "true",
			Message:   "Success",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"music": &music,
			}})
	}

}

func GetMusic(c *gin.Context) {
	var music MusicsDetail
	musicId := c.Param("id")

	opts := mongoOptions.FindOne().SetProjection(bson.D{{Key: "deleted", Value: 0}})

	err := configs.MongoDbCollection(client, "musics", "music").FindOne(context.TODO(), bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{
				Key:   "id",
				Value: musicId,
			}}, ignoreDeleted,
		}}}, opts).Decode(&music)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Success:   "false",
			Message:   "Not found",
			ErrorCode: 1,
			Data:      "null",
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Success:   "true",
			Message:   "Success",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"music": &music,
			}})
	}
}

func CreateMusic(c *gin.Context) {
	var music MusicsDetail
	musicId := Id(time.Now().String())
	// log.Println(musicId)

	err := c.BindJSON(&music)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success:   "false",
			Message:   "Bad Request",
			ErrorCode: 3,
			Data: map[string]interface{}{
				"error": err.Error(),
			}})
	}

	newMusic := MusicsDetail{
		Id:             musicId,
		Singer:         music.Singer,
		SongName:       music.SongName,
		Album:          music.Album,
		Release:        music.Release,
		RecordingLabel: music.RecordingLabel,
	}

	result, err := configs.MongoDbCollection(client, "musics", "music").InsertOne(context.TODO(), newMusic)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success:   "false",
			Message:   "Internal Server Error",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"error": err.Error(),
			}})
		configs.MongoDbLogger()
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, Response{
			Success:   "true",
			Message:   "Success",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"music":  &newMusic,
				"result": result,
			}})
	}
}

func DeleteMusic(c *gin.Context) {
	musicId := c.Param("id")
	result, err := configs.MongoDbCollection(client, "musics", "music").UpdateOne(context.TODO(), bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{
				Key:   "id",
				Value: musicId,
			}}, ignoreDeleted,
		}}}, bson.D{{
		Key: "$set",
		Value: bson.D{{
			Key:   "deleted",
			Value: true,
		}}}})
	// if result
	fmt.Println(result)
	if err != nil {
		configs.MongoDbLogger()
		log.Println(err)
		c.JSON(http.StatusInternalServerError, Response{
			Success:   "false",
			Message:   "Internal Server Error",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"error": err.Error(),
			}})
	} else if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, Response{
			Success:   "false",
			ErrorCode: 1,
			Message:   "Music not found",
			Data: map[string]interface{}{
				"musicId": musicId,
			}})
		// } else if result.ModifiedCount == 0 {
		// 	c.JSON(http.StatusBadRequest, Response{
		// 		Success:   "false",
		// 		ErrorCode: 3,
		// 		Message:   "Bad Request",
		// 		Data:      "null",
		// 	})
		// }

	} else {
		c.JSON(http.StatusOK, Response{
			Success:   "true",
			ErrorCode: 0,
			Message:   "Success",
			Data: map[string]interface{}{
				"musicId": musicId,
			}})
	}
}
func UpdateMusic(c *gin.Context) {
	var music MusicsDetail

	musicId := c.Param("id")

	err := c.BindJSON(&music)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Success:   "false",
			Message:   "Bad Request",
			ErrorCode: 3,
			Data: map[string]interface{}{
				"error": err.Error(),
			}})
	}

	newMusicData := bson.D{{Key: "singer", Value: music.Singer}, {Key: "songName", Value: music.SongName}, {Key: "album", Value: music.Album}, {Key: "release", Value: music.Release}, {Key: "recordingLabel", Value: music.RecordingLabel}}

	result, err := configs.MongoDbCollection(client, "musics", "music").UpdateOne(context.TODO(), bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{
				Key:   "id",
				Value: musicId,
			}}, ignoreDeleted,
		}}}, bson.D{{
		Key:   "$set",
		Value: newMusicData}})
	// if result
	// if err != nil {
	// 	configs.MongoDbLogger()
	// 	log.Fatalln(err)
	// }

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Success:   "false",
			Message:   "Internal Server Error",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"error": err.Error(),
			}})
		configs.MongoDbLogger()
		log.Println(err)
	} else if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, Response{
			Success:   "false",
			ErrorCode: 1,
			Message:   "Music not found",
			Data: map[string]interface{}{
				"musicId": musicId,
			}})
	} else {
		c.JSON(http.StatusCreated, Response{
			Success:   "true",
			Message:   "Success",
			ErrorCode: 0,
			Data: map[string]interface{}{
				"musicId": musicId,
			}})
	}
	// c.JSON(http.StatusOK, result)
}
