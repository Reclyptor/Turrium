package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func EnableCors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:  []string{"https://turrium.com", "http://localhost:3000"},
		AllowMethods:  []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
}