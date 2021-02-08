package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/repository"
)

func GetAlbums(c *gin.Context) {
	c.JSON(200, repository.GetAlbums(bson.M{"hidden": false}, 15 * time.Minute))
}