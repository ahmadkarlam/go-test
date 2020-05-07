package main

import "github.com/gin-gonic/gin"

type Container struct {
	numberOfBall int
	verified     bool
}

func main() {
	r := gin.Default()

	r.GET("/containers", func(ctx *gin.Context) {

	})
}
