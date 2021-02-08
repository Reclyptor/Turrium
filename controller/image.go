package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/service"
)

func GetImages(c *gin.Context) {
	c.JSON(200, service.GetImages(bson.M{}, 15 * time.Minute)[:100])
}