package main

import (
	"log"

	"github.com/allefts/minedle/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.Default())
	router.GET("/", routes.GetItemRoute)

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowAllOrigins = true
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowHeaders("Access-Control-Allow-Origin: *")
	// corsConfig.AddAllowMethods("GET")
	// router.Use(cors.New(corsConfig))

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:5173/"},
	// 	AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
	// 	AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
