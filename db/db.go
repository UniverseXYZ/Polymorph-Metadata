package db

import (
	"context"
	"os"
	"sync"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var once sync.Once
var instance *mongo.Client

// GetDbConnection returns an instance of a mongo db client. This is a singleton pattern in order to have only one alive connection to the database.
//
// If no connection exists, it will connect to database.
//
// If connection exists, it will return the instance of the database
func GetDbConnection() *mongo.Client {
	once.Do(func() {
		client := ConnectToDb()
		checkConnectionAndRestore(client)
		instance = client
	})

	checkConnectionAndRestore(instance)
	return instance
}

// ConnectToDb retrieves db config from .env and tries to conenct to the database.
func ConnectToDb() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Errorln("Error loading .env file")
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	dbUrl := os.Getenv("DB_URL")

	if username == "" {
		log.Errorln("Missing username in .env")
	}
	if password == "" {
		log.Errorln("Missing password in .env")
	}
	if dbUrl == "" {
		log.Errorln("Missing db url in .env")
	}

	connectionStr := "mongodb+srv://" + username + ":" + password + "@" + dbUrl + "?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionStr))
	if err != nil {
		log.Errorln(err)
	}

	return client
}

// checkConnectionAndRestore ping the client and it throws and error, it tries to reconnect.
func checkConnectionAndRestore(client *mongo.Client) {
	err := client.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Errorln(err)
		newClient := ConnectToDb()
		client = newClient
	}
}

// GetMongoDbCollection accepts dbName and collectionname and returns an instance of the specified collection.
func GetMongoDbCollection(DbName string, CollectionName string) (*mongo.Collection, error) {
	client := GetDbConnection()

	collection := client.Database(DbName).Collection(CollectionName)
	return collection, nil
}
