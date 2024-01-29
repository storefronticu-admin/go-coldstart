package firebase

import (
	"fmt"

	"storefront.icu/go-coldstart/logger"
)

// UpdateDocument updates a document in Firestore
func UpdateDocument(collection string, documentID string, data map[string]interface{}) error {
	_, err := firestoreClient.Collection(collection).Doc(documentID).Set(ctx, data)
	if err != nil {
		logger.Error("An error has occurred: %v", err)
		return fmt.Errorf("Failed updating document")
	}

	logger.Success("Success writing to Firestore")
	return nil
}

// ReadSpecificDocument reads a specific document from Firestore based on the provided ID
func ReadDocument(collection string, documentID string) (map[string]interface{}, error) {
	dsnap, err := firestoreClient.Collection(collection).Doc(documentID).Get(ctx)
	if err != nil {
		logger.Error("Failed to retrieve document: %v", err)
		return nil, fmt.Errorf("Failed to retrieve document")
	}

	return dsnap.Data(), nil
}


