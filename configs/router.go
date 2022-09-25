package configs

import (
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(router *gin.Engine) {
	mode := RouterConf()
	if mode == "release" {
		path, _ := os.Getwd()
		logPath := path + "/log/" + time.Now().Format("01-02-2006") + ".log"
		accessLog, _ := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		gin.DefaultWriter = io.MultiWriter(accessLog)
	}
	router.Use(gin.Logger())
}

func Mode() {
	mode := RouterConf()
	gin.SetMode(mode)
}
