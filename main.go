package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/firestore/apiv1/firestorepb"
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
	router.LoadHTMLGlob("resources/*.templ.html")
	router.Static("/static", "resources/static")
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


// repoinfoPost is a placeholder for the actual implementation of the POST method for repoinfo.
func repoinfoPost(c *gin.Context) {
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		errorLog.Printf("Failed adding Lovelace: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %v", err)})
		return
	}

	successLog.Println("Success writing to Firestore")
	c.JSON(http.StatusOK, gin.H{"message": "POST request handled"})
	debugLog.Println("repoinfoPost is not implemented yet")
}

// repoinfoPut is a placeholder for the actual implementation of the PUT method for repoinfo.
func repoinfoPut(c *gin.Context) {
	// TODO: Implement the PUT method logic here.
	debugLog.Println("repoinfoPut is not implemented yet")
}

// repoinfoGet is a placeholder for the actual implementation of the GET method for repoinfo.
func repoinfoGet(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		collection := client.Collection("users")

		aggregationQuery := collection.NewAggregationQuery().WithCount("all")
		results, err := aggregationQuery.Get(ctx)
		if err != nil {
			errorLog.Printf("Error querying: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %v", err)})
			return
		}

		count, ok := results["all"]
		if !ok {
			errorLog.Println("Firestore: couldn't get alias for COUNT from results")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving count"})
			return
		}

		countValue := count.(*firestorepb.Value)
		resultString := fmt.Sprintf("Number of results from query: %d\n", countValue.GetIntegerValue())
		successLog.Println(resultString)

		// Convert the int64 to a string before writing it to the response
		c.JSON(http.StatusOK, gin.H{"count": strconv.FormatInt(countValue.GetIntegerValue(), 10)})
		debugLog.Println("repoinfoHead is not implemented yet")
		return
	}

	dsnap, err := client.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve document"})
		return
	}

	m := dsnap.Data()

	if m != nil {
		c.JSON(http.StatusOK, m)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "GET request handled, but no data available"})
	debugLog.Println("repoinfoGet is not implemented yet")
}

