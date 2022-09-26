package controller

import (
	"context"
	"log"
	"net/http"
	"time"

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
	collection, err := configs.MongoDbCollection(client, "musics", "music").Find(context.TODO(), bson.D{})
	err = collection.All(context.TODO(), &music)
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, &music)

}

func GetMusic(c *gin.Context) {
	var music MusicsDetail

	musicId := c.Param("id")
	client, _ := configs.MongoDb()

	err := configs.MongoDbCollection(client, "musics", "music").FindOne(context.TODO(), bson.M{"id": musicId}).Decode(&music)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, &music)
}

func CreateMusic(c *gin.Context) {
	var music MusicsDetail

	timeNow := time.Now().String()

	musicId := configs.Id(timeNow)
	// log.Println(musicId)
	c.BindJSON(&music)

	newMusic := MusicsDetail{
		Id:             musicId,
		Singer:         music.Singer,
		SongName:       music.SongName,
		Album:          music.Album,
		Release:        music.Release,
		RecordingLabel: music.RecordingLabel,
	}

	client, _ := configs.MongoDb()

	result, _ := configs.MongoDbCollection(client, "musics", "music").InsertOne(context.TODO(), newMusic)
	// log.Println(result)
	c.JSON(http.StatusCreated, result)

}
