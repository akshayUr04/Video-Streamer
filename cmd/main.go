package main

import (
	"github.com/akshayUr04/video-streaming/pkg/streamer"
	"github.com/akshayUr04/video-streaming/pkg/uploder"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("upload", uploder.Upload)
	r.GET("play/:id/:playList", streamer.Streamer)
	r.Run("localhost:8080")
}
