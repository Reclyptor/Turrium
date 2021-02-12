package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"turrium/controller"
	"turrium/env"
	"turrium/middleware"
)

func main() {
	env.Verify()
	router := gin.Default()

	router.GET("/info", controller.GetInfo)

	router.NoRoute(controller.GetReact)
	router.GET("/login", func(c *gin.Context) {c.Redirect(http.StatusTemporaryRedirect, "/")})

	api := router.Group("/api")
	api.Use(middleware.EnableCors())
	api.Use(middleware.VerifyTokens())
	api.Use(middleware.VerifyUser())
	v1 := api.Group("/v1")
	v1.GET("/albums", controller.GetAlbums)
	v1.GET("/images", controller.GetImages)
	v1.GET("/videos", controller.GetVideos)

	log.Fatal(router.Run())
}