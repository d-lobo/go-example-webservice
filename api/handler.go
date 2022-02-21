package api

import (
	"example/webservice/service"
	"example/webservice/service/model"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAllAlbums())
}

func addAlbum(c *gin.Context) {
	var alb *model.AlbumDTO
	err := c.BindJSON(&alb)
	if err != nil {
		log.Errorf("cannot bind payload to model: %v", err)
		c.IndentedJSON(http.StatusBadRequest, "{msg: error}")
		return
	}

	a, err := service.CreateNewAlbum(alb)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "{msg: error}")
		return
	}

	c.IndentedJSON(http.StatusCreated, a)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	a, err := service.GetAlbumById(id)
	if err != nil {
		log.Errorf("cant retrieve album with id: %v, %v", id, err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, a)
}
