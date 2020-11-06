package main

import (
	"log"

	"github.com/mastanca/go-api-template/cmd/web/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	pingHandler := handlers.NewPingHandlerImpl()

	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", pingHandler.Handle)
	}
	if err := router.Run(":8080"); err != nil {
		log.Fatal("error initializing server")
	}
}
