package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"turrium/structs"
)

func GetInfo(c *gin.Context) {
	var JSON map[string]interface{}

	data, _ := ioutil.ReadFile("info.json")
	err := json.Unmarshal(data, &JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.Status{
			Code:   http.StatusNotFound,
			Status: "NOT_FOUND",
			Reason: "No information is available.",
		})
		return
	}

	c.JSON(200, JSON)
}