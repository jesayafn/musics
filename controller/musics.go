package controller

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jessie-txt/musics/configs"
	"go.mongodb.org/mongo-driver/bson"
)

type Musics []struct {
	Id             string   `json:"id"`
	Singer         []string `json:"singer"`
	SongName       string   `json:"songName"`
	Album          string   `json:"album"`
	Release        string   `json:"release"`
	RecordingLabel string   `json:"recordingLabel"`
}

func GetMusics(c *gin.Context) {
	var music Musics
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
