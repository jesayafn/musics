package main

import (
	// "context"
	// json "encoding/json"
	// "log"
	// http "net/http"
	// "os"

	gin "github.com/gin-gonic/gin"
	// configs "github.com/jessie-txt/web-backend-example/configs"
	configs "github.com/jessie-txt/musics/configs"
	controller "github.com/jessie-txt/musics/controller"
	// "golang.org/x/text/date"
)

func main() {
	// configs.MongoDb()
	configs.Mode()
	router := gin.New()
	configs.Logger(router)
	router.GET("/musics", controller.GetMusics)
	router.GET("/musics/:id", controller.GetMusic)
	router.POST("/musics", controller.CreateMusic)
	router.Run(":5678")

}

// func parseJson(file string) (*Musics, error) {
// 	jsonFile, err := os.ReadFile(file)
// 	if err != nil {
// 		log.Printf("jsonFile.Get err\n %v", err)
// 	}

// 	jsonData := &Musics{}

// 	err = json.Unmarshal(jsonFile, jsonData)

// 	if err != nil {
// 		log.Printf("Unmarshal: %v", err)
// 	}

// 	return jsonData, err
// }

// func getMusics(c *gin.Context) {
// 	var filePath string = "music.json"
// 	musicsData, err := parseJson(filePath)
// 	if err != nil {
// 		log.Printf("%v", err)
// 	}
// 	c.IndentedJSON(http.StatusOK, musicsData)
// }

// func postMusic(c *gin.Context) {
// 	var filePath string = "music.json"
// 	musicsData, err := parseJson(filePath)
// 	if err != nil {
// 		log.Printf("%v", err)
// 	}
// 	var newMusicsData Musics

// 	newMusicsData = append(*musicsData, newMusicsData...)

// 	dataBytes, err := json.Marshal(newMusicsData)
// 	if err != nil {
// 		log.Printf("err %v", err)
// 	}

// 	err = os.WriteFile(filePath, dataBytes, 0644)
// 	if err != nil {
// 		log.Printf("err %v", err)
// 	}
// 	c.IndentedJSON(http.StatusCreated, newMusicsData)
// }
