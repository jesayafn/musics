package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessie-txt/musics/configs"
	"go.mongodb.org/mongo-driver/bson"
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

func GetMusics(c *gin.Context) {
	var music MusicOverview

	client, _ := configs.MongoDb()
	collection, err := client.Database("musics").Collection("music").Find(context.TODO(), bson.D{})
	err = collection.All(context.TODO(), &music)
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println(music)
	// if err := c.BindJSON(music); err != nil {
	// 	return
	// }
	c.JSON(http.StatusOK, &music)

}

func GetMusic(c *gin.Context) {
	var music MusicsDetail

	musicId := c.Param("id")
	// objId, _ := primitive.ObjectIDFromHex(musicId)
	client, _ := configs.MongoDb()
	err := client.Database("musics").Collection("music").FindOne(context.TODO(), bson.M{"id": musicId}).Decode(&music)
	if err != nil {
		log.Fatalln(err)
	}
	// log.Println(&music)
	// if err := c.BindJSON(music); err != nil {
	// 	return
	// }
	c.JSON(http.StatusOK, &music)
}
