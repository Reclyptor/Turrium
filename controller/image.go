package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/repository"
)

func GetImages(c *gin.Context) {
	c.JSON(200, repository.GetImages(bson.M{}, 15 * time.Minute))
}