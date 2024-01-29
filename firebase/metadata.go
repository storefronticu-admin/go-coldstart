package firebase

import (
	"fmt"

	"storefront.icu/go-gin-ledger-server/logger"
)

// UpdateDocument updates a document in Firestore
func CountDocuments(collection string, documentID string, data map[string]interface{}) error {
	_, err := firestoreClient.Collection(collection).Doc(documentID).Set(ctx, data)
	if err != nil {
		logger.Error("An error has occurred: %v", err)
		return fmt.Errorf("Failed updating document")
	}

	logger.Success("Success writing to Firestore")
	return nil
}
