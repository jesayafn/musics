package main

import (
	// "context"
	// json "encoding/json"
	// "log"
	// http "net/http"
	// "os"

	gin "github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	// configs "github.com/jessie-txt/web-backend-example/configs"
	configs "github.com/jessie-txt/musics/configs"
	controller "github.com/jessie-txt/musics/controller"
	// "golang.org/x/text/date"
)

func init() {
	prometheus.MustRegister(controller.TotalRequests)
}

func main() {
	// configs.MongoDb()

	configs.GinMode()
	router := gin.New()
	configs.GinLogger(router)
	// metrics := ginmetrics.GetMonitor()
	// metrics.SetMetricPath("/metrics")
	// metrics.Use(router)

	router.GET("/musics", controller.GetMusics)
	router.GET("/musics/:id", controller.GetMusic)
	router.POST("/musics", controller.CreateMusic)
	router.DELETE("/musics/:id", controller.DeleteMusic)
	router.PATCH("/musics/:id", controller.UpdateMusic)
	router.GET("/metrics", prometheusHandler())
	router.Run(":5678")
	// prometheus.Register(controller.TotalRequests)

}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
