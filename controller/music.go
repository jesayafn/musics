package controller

import (
	"context"
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

var (
	timeNow     = time.Now().String()
	client, err = configs.MongoDb()
)

func GetMusics(c *gin.Context) {
	var music MusicOverview

	opts := mongoOptions.Find().SetProjection(bson.D{{Key: "deleted", Value: 0}})
	collection, err := configs.MongoDbCollection(client, "musics", "music").Find(context.TODO(),
		bson.D{{Key: "$or",
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
		}}, opts)

	if err != nil {
		configs.MongoDbLogger()
		log.Fatalln(err)
	}

	err = collection.All(context.TODO(), &music)
	if err != nil {
		configs.MongoDbLogger()
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, &music)

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
			}}, bson.D{{Key: "$or",
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
			}},
		}}}, opts).Decode(&music)
	if err != nil {
		c.JSON(http.StatusNotFound, "yah error, wwkkw")
	} else {
		c.JSON(http.StatusOK, &music)
	}
}

func CreateMusic(c *gin.Context) {
	var music MusicsDetail
	musicId := Id(time.Now().String())
	// log.Println(musicId)

	err := c.BindJSON(&music)

	if err != nil {
		log.Fatalln(err)
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
		configs.MongoDbLogger()
		log.Fatalln(err)
	}

	// log.Println(result)
	c.JSON(http.StatusCreated, result)

}

func DeleteMusic(c *gin.Context) {
	musicId := c.Param("id")
	result, err := configs.MongoDbCollection(client, "musics", "music").UpdateOne(context.TODO(), bson.D{{
		Key: "$and",
		Value: bson.A{
			bson.D{{
				Key:   "id",
				Value: musicId,
			}}, bson.D{{Key: "$or",
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
			}},
		}}}, bson.D{{Key: "$set", Value: bson.D{{Key: "deleted", Value: true}}}})
	// if result
	if err != nil {
		configs.MongoDbLogger()
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, result)
}
