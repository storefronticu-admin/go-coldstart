package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//  Base path

func RepoinfoGet(c *gin.Context) {
	// id := c.Query("id")
	// if id == "" {
	// 	logger.Error("Failed getting specific repo info")
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Repository identifier is missing to retrieve document"})
	// 	return
	// }

	// dsnap, err := firestoreClient.Collection("users").Doc(id).Get(ctx)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve document"})
	// 	return
	// }

	// m := dsnap.Data()

	// if m != nil {
	// 	c.JSON(http.StatusOK, m)
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "GET request handled, but no data available"})
}

func RepoinfoPost(c *gin.Context) {
	// 	repository := c.Query("repo")
	// _, _, err := firestoreClient.Collection("users").Add(ctx, map[string]interface{}{
	// 	"repository": repository,
	// 	"last":  "Lovelace",
	// 	"born":  1815,
	// })
	// if err != nil {
	// 	logger.Error("Failed adding Lovelace: %v", err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %v", err)})
	// 	return
	// }

	// logger.Success("Success writing to Firestore")
	c.JSON(http.StatusOK, gin.H{"message": "POST request handled"})
}


func RepoinfoPut(c *gin.Context) {
// // City represents a city.
// type Repository struct {
//        Index       string   `firestore:"index,omitempty"`
// }

// 		id := c.Query("repo")
		


// 		  repository := Repository{
//                 Index:    id,
//         }

// 		   _, err := firestoreClient.Collection("users").Doc(id).Set(ctx, repository)
//         if err != nil {
//                 // Handle any errors in an appropriate way, such as returning them.
//                  logger.Error("An error has occurred: %s", err)
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error: %v", err)})
//         }

// 	logger.Success("Success writing to Firestore")
	c.JSON(http.StatusOK, gin.H{"message": "POST request handled"})
}
