package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"turrium/controller"
	"turrium/env"
	"turrium/middleware"
)

func main() {
	env.Verify()
	router := gin.Default()
	router.Use(middleware.EnableCors())
	router.Use(middleware.ServeReact())
	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/albums", controller.GetAlbums)
	v1.GET("/images", controller.GetImages)
	v1.GET("/videos", controller.GetVideos)
	log.Fatal(router.Run())
}