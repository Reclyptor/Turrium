package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/repository"
	"turrium/structs"
)

func GetVideos(c *gin.Context) {
	var pagination structs.Pagination
	err := c.Bind(&pagination)
	if err != nil {
		pagination = structs.Pagination{Page: 1, Size: 50}
	}
	c.JSON(200, repository.GetVideos(bson.M{"filename": bson.M{"$regex": "\\.mp4$"}}, pagination, 15 * time.Minute))
}