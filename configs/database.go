package configs

import (
	"context"
	"log"
	"time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

func MongoDb() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	uriDb, authSourceDb, usernameDb, passwordUserDb := DatabaseConf()
	credential := mongoOptions.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    authSourceDb,
		Username:      usernameDb,
		Password:      passwordUserDb,
	}
	clientOptions := mongoOptions.Client().ApplyURI(uriDb).SetAuth(credential)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	// log.Println("Connected to MongoDB")

	return client, err

}

func MongoDbCollection(client *mongo.Client, database string, collectionName string) (collectionResult *mongo.Collection) {
	collectionResult = client.Database(database).Collection(collectionName)
	return collectionResult
}
