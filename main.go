package main

import (
	"log"

	"github.com/allefts/minedle/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", routes.GetItemRoute)
	router.GET("/", routes.GetRandomItemRoute)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
