package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/repository"
)

func GetVideos(c *gin.Context) {
	c.JSON(200, repository.GetVideos(bson.M{"filename": bson.M{"$regex": "\\.mp4$"}}, 15 * time.Minute))
}