package configs

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GinLogger(router *gin.Engine) {
	_, env := RouterConf()
	if env != "container" {
		path, _ := os.Getwd()
		logPath := path + "/log/" + time.Now().Format("01-02-2006") + ".log"
		accessLog, _ := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		gin.DefaultWriter = io.MultiWriter(accessLog)
	}
	router.Use(gin.Logger())
}

func GinMode() {
	mode, _ := RouterConf()
	gin.SetMode(mode)
}
