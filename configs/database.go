package configs

import (
	"context"
	"log"
	"os"
	"time"

	mongo "go.mongodb.org/mongo-driver/mongo"
	mongoOptions "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
)

func MongoDb() (*mongo.Client, error) {
	uriDb, authSourceDb, usernameDb, passwordUserDb := DatabaseConf()
	credential := mongoOptions.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    authSourceDb,
		Username:      usernameDb,
		Password:      passwordUserDb,
	}
	clientOptions := mongoOptions.Client().ApplyURI(uriDb).SetAuth(credential)
	// fmt.Println("Ini")
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		MongoDbLogger()
		log.Fatalln(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		MongoDbLogger()
		log.Fatalln(err)
	}

	// log.Println("Connected to MongoDB")

	return client, err

}

func MongoDbCollection(client *mongo.Client, database string, collectionName string) (collectionResult *mongo.Collection) {
	collectionResult = client.Database(database).Collection(collectionName)
	return collectionResult
}
func MongoDbLogger() {
	_, env := RouterConf()
	if env != "container" {
		path, _ := os.Getwd()
		logPath := path + "/log/" + time.Now().Format("01-02-2006") + ".log"
		accessLog, _ := os.OpenFile(logPath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		log.SetOutput(accessLog)
	}
	log.SetPrefix("[MONGODB] ")
}
