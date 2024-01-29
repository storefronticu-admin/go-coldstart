package firebase

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"storefront.icu/go-coldstart/logger"
)

var (
	ctx            context.Context
	firestoreClient *firestore.Client
)

// init initializes the Firestore client and logs success or failure.
func init() {
	// Use a service account
	ctx = context.Background()
	sa := option.WithCredentialsFile("./serviceAccount.json")

	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		logger.Error("Error initializing Firebase app: %v", err)
	}
	logger.Success("Firebase app initialized")

	firestoreClient, err = app.Firestore(ctx)
	if err != nil {
		logger.Error("Error initializing Firestore client: %v", err)
	}

	logger.Success("Firestore client initialized")
}

// QueryDocuments queries documents in Firestore based on the provided parameters.
// It returns a slice of maps representing the queried documents or an error if the operation fails.
//
// Parameters:
// - collection: Firestore collection name.
// - queryParam: Query parameter for filtering documents.
// - queryValue: Value to match for the query parameter.
//
// Returns:
// - []map[string]interface{}: Queried documents.
// - error: Error if the operation fails.
func QueryDocuments(collection string, queryParam string, queryValue interface{}) ([]map[string]interface{}, error) {
	query := firestoreClient.Collection(collection).Where(queryParam, "==", queryValue)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {
		logger.Error("Failed to query documents: %v", err)
		return nil, fmt.Errorf("Failed to query documents")
	}

	var result []map[string]interface{}
	for _, doc := range docs {
		result = append(result, doc.Data())
	}

	return result, nil
}

// AddDocument adds a new document to Firestore.
// It takes a collection name and a map of data to be added and returns an error if the operation fails.
//
// Parameters:
// - collection: Firestore collection name.
// - data: Data to be added to Firestore.
//
// Returns:
// - error: Error if the operation fails.
func AddDocument(collection string, data map[string]interface{}) error {
	_, _, err := firestoreClient.Collection(collection).Add(ctx, data)
	if err != nil {
		logger.Error("Failed adding document: %v", err)
		return fmt.Errorf("Failed adding document")
	}

	logger.Success("Success writing to Firestore")
	return nil
}
