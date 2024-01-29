package main

import (
	"github.com/gin-gonic/gin"
)

// Count path

func RepoinfoCountGet(c *gin.Context) {
		// collection := firestoreClient.Collection("users")
		// aggregationQuery := collection.NewAggregationQuery().WithCount("all")
		// results, err := aggregationQuery.Get(ctx)
		// if err != nil {
		// 	logger.Error("Error querying: %v", err)
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %v", err)})
		// 	return
		// }

		// count, ok := results["all"]
		// if !ok {
		// 	logger.Error("Firestore: couldn't get alias for COUNT from results")
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving count"})
		// 	return
		// }

		// countValue := count.(*firestorepb.Value)
		// resultString := fmt.Sprintf("Number of results from query: %d\n", countValue.GetIntegerValue())
		// logger.Success(resultString)

		// Convert the int64 to a string before writing it to the response
		// c.JSON(http.StatusOK, gin.H{"count": strconv.FormatInt(countValue.GetIntegerValue(), 10)})
		// return
}
