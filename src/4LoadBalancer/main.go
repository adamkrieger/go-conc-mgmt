package main

import (
	"./balancer"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

var (
	rrBalancer balancer.RoundRobin
)

func main() {
	rrBalancer = balancer.StartRoundRobin([]string{"localhost:9000", "localhost:9001", "localhost:9002"})

	ngn := gin.Default()
	ngn.Any("/*path", proxy())

	_ = ngn.Run(":8000")
}

func proxy() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		director := func(rqst *http.Request) {
			rqst.URL.Scheme = "http"

			nextHost := <-rrBalancer.NextHost()
			rqst.URL.Host = nextHost

			rqst.Header["my-header"] = []string{ctx.Request.Header.Get("my-header")}
			// Golang camelcases headers
			delete(rqst.Header, "My-Header")
		}

		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
