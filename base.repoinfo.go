package main

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/gin-gonic/gin"
)

//  Base path

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

func repoinfoPut(c *gin.Context) {
	// TODO: Implement the PUT method logic here.
	debugLog.Println("repoinfoPut is not implemented yet")
}
