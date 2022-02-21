package api

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", addAlbum)
	router.Run("localhost:8080")
}
