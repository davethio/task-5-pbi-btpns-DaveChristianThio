package main

import (
	"fmt"
	"log"

	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/database"
	"github.com/davethio/task-5-pbi-btpns-DaveChristianThio/router"
	"github.com/gin-gonic/gin"
)

func main() {

	// Inisialize Gin router
	r := gin.Default()
	database.ConnectDatabase()

	router.SetupRouter(r)

	// Port to use
	port := 8080

	// Run server Gin
	serverAddr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is running on port %d...\n", port)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	r.Run()

}
