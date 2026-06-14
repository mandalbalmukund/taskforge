package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"taskforge/internal/api"
)

func main() {

	router := gin.Default()

	// initialize repositories
	// initialize services
	// initialize handlers

	api.RegisterRoutes(
		router,
		workflowHandler,
		executionHandler,
		approvalHandler,
		cancelHandler,
	)

	log.Fatal(
		router.Run(":8080"),
	)
}