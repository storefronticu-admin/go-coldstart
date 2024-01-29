package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//  Base path

func RepoinfoGet(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		errorLog.Printf("Failed getting specific repo info")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository identifier is missing to retrieve document"})
		return
	}

	dsnap, err := firestore_client.Collection("users").Doc(id).Get(ctx)
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
}

func RepoinfoPost(c *gin.Context) {
	_, _, err := firestore_client.Collection("users").Add(ctx, map[string]interface{}{
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
}

