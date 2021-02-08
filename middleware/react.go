package middleware

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func ServeReact() gin.HandlerFunc {
	return static.Serve("/", static.LocalFile("./ui", true))
}