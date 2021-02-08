package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"time"
	"turrium/repository"
)

func Login(c *gin.Context) {

}

func Logout(c *gin.Context) {
	c.JSON(200, repository.GetImages(bson.M{}, 15 * time.Minute)[:100])
}
