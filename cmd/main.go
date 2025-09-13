package main

import (
	"log"
	"money-tracker/config"
	"money-tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(":8081"); err != nil {
		log.Fatalf("❌ Server failed to start: %v", err)
	}
}
