package connections

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateClient() *mongo.Client{
	client, _ := mongo.NewClient(options.Client().ApplyURI("mongodb://mongoadmin:secret@172.17.0.3.:27017"))

	return client
}