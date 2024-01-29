package main

import (
	"fmt"
	"os"
	"runtime"

	"storefront.icu/go-gin-ledger-server/logger"

	"github.com/gin-gonic/gin"
)

// main is the entry point of the application.
func main() {
	ConfigRuntime()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers starts statsWorker by goroutine.

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(rateLimit, gin.Recovery())
	router.GET("/", index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Add the repoinfo API group
	repoinfoGroup := router.Group("/repoinfo")
	{
		repoinfoGroup.GET("/", RepoinfoGet) // Get repository information
		repoinfoGroup.POST("/", RepoinfoPost) // Post repository information
		repoinfoGroup.GET("/count", RepoinfoCountGet) // Get count of all repositories
	}

	// Log that the server is starting
	logger.Success("Server is starting on http://localhost:%s...", port)

	// Use a goroutine to run the server
	go func() {
		addr := ":" + port
		if err := router.Run(addr); err != nil {
			logger.Error("error: %s", err)
		}
	}()

	// Log that the server has started successfully
	logger.Success("Server is now running on http://localhost:%s", port)

	// The main function will not exit immediately
	select {}
}
