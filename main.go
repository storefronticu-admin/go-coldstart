package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	ConfigRuntime()
	StartWorkers()
	StartGin()
}

// ConfigRuntime sets the number of operating system threads.
func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

// StartWorkers start starsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

// StartGin starts gin web server with setting router.
func StartGin() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(rateLimit, gin.Recovery())
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
	router.GET("/", index)
	router.GET("/room/:roomid", roomGET)
	router.POST("/room-post/:roomid", roomPOST)
	router.GET("/stream/:roomid", streamRoom)

 port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Add the myinfo API group
    myinfoGroup := router.Group("/myinfo")
    {
        myinfoGroup.POST("/", myinfoPost)
        myinfoGroup.PUT("/", myinfoPut)
        myinfoGroup.GET("/", myinfoGet)
        myinfoGroup.HEAD("/", myinfoHead)
    }

 // Log that the server is starting
    log.Printf("Server is starting on http://localhost:%s...", port)

    // Use a goroutine to run the server
    go func() {
        addr := ":" + port
        if err := router.Run(addr); err != nil {
            log.Panicf("error: %s", err)
        }
    }()

    // Log that the server has started successfully
    log.Printf("Server is now running on http://localhost:%s", port)

    // The main function will not exit immediately
    select {}

}


// Handler functions for myinfo API group
func myinfoPost(c *gin.Context) {
    // Handle POST request for /myinfo
    c.JSON(200, gin.H{"message": "myinfo POST request"})
}

func myinfoPut(c *gin.Context) {
    // Handle PUT request for /myinfo
    c.JSON(200, gin.H{"message": "myinfo PUT request"})
}

func myinfoGet(c *gin.Context) {
    // Handle GET request for /myinfo
    c.JSON(200, gin.H{"message": "myinfo GET request"})
}

func myinfoHead(c *gin.Context) {
    // Handle HEAD request for /myinfo
    c.Status(200)
}