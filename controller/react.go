package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"turrium/env"
)

func GetReact(c *gin.Context) {
	remote, _ := url.Parse(env.TURRIUM_UI_URL)
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}