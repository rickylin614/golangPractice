package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Any("/", func(ctx *gin.Context) {
		remote, err := url.Parse("https://lv-admin-uat.paradise-soft.com.tw/")
		if err != nil {
			return
		}
		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
		return
	})

	r.Run()
}
