package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/service"
)

func GetAlbums(c *gin.Context) {
	c.JSON(200, service.GetAlbums(bson.M{"hidden": false}, 15 * time.Minute))
}