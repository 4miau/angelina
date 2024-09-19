package database

import "go.mongodb.org/mongo-driver/mongo"

var (
	mongoClient **mongo.Client // this is a memory address
)

func SetClient(m **mongo.Client) { mongoClient = m }

func GetClient() *mongo.Client { return *mongoClient }
