package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"turrium/controller"
)

func main() {
	router := gin.Default()

	router.Use(
		static.Serve("/", static.LocalFile("./ui", true)),
		static.Serve("/signin", static.LocalFile("./ui", true)),
	)

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://turrium.com, http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	api := router.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/albums", controller.GetAlbums)
	v1.GET("/images", controller.GetImages)
	v1.GET("/videos", controller.GetVideos)

	log.Fatal(router.Run())
}