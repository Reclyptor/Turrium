package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/service"
)

func GetVideos(c *gin.Context) {
	c.JSON(200, service.GetVideos(bson.M{"filename": bson.M{"$regex": "\\.mp4$"}}, 15 * time.Minute))
}