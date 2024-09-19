package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(uri string) {
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	defer func() {
		if err := cli.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	SetClient(&cli)
}

func ConnectDB() {
	if checkAngelinaDB() != true {
		createAngelinaDB()
	}
}
