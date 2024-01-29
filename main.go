package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// Create a new colorized logger
	var infoLog = log.New(os.Stdout, color.GreenString("INFO: "), log.Ldate|log.Ltime)
	var warningLog = log.New(os.Stdout, color.YellowString("WARNING: "), log.Ldate|log.Ltime|log.Lshortfile)
	var errorLog = log.New(os.Stderr, color.RedString("ERROR: "), log.Ldate|log.Ltime|log.Lshortfile)
	var successLog = log.New(os.Stdout, color.MagentaString("SUCCESS: "), log.Ldate|log.Ltime)
	var debugLog = log.New(os.Stdout, color.BlueString("DEBUG: "), log.Ldate|log.Ltime)

var (
	ctx    context.Context
	client *firestore.Client
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

	// Use a service account
	ctx = context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		errorLog.Fatalf("Error initializing Firebase app: %v", err)
	}
	successLog.Println("Firebase app initialized")

	client, err = app.Firestore(ctx)
	if err != nil {
		errorLog.Fatalf("Error initializing Firestore client: %v", err)
	}
	successLog.Println("Firestore client initialized")
}

// StartWorkers start statsWorker by goroutine.
func StartWorkers() {
	go statsWorker()
}

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
		repoinfoGroup.POST("/", repoinfoPost)
		repoinfoGroup.PUT("/", repoinfoPut)
		repoinfoGroup.GET("/", repoinfoGet)
		// 
		repoinfoGroup.GET("/count", repoinfoCountGet)
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
