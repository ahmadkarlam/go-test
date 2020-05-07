package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	collection := NewCollection()
	r := gin.Default()

	r.POST("/fill", func(ctx *gin.Context) {
		if collection.isThereAFullContainer() {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "there are already full containers",
				"data":    collection.containers,
			})
			return
		}

		rand.Seed(time.Now().UTC().UnixNano())

		key := rand.Intn(math.MaxInt8)
		collection.fill(key)

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("success fill one ball to container %d", key),
			"data":    collection.containers,
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
