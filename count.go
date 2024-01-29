package main

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/gin-gonic/gin"
)

// Count path

func RepoinfoCountGet(c *gin.Context) {
		collection := firestore_client.Collection("users")

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
		return
}
