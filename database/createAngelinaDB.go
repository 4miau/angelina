package database

import (
	"context"
	"fmt"
)

func checkAngelinaDB() bool {
	client := GetClient()
	databases, err := client.ListDatabaseNames(context.TODO(), "")
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	fmt.Printf("%q\n", databases)

	for _, val := range databases {
		if val == "angelina" {
			return true
		}
	}

	return false
}

func createAngelinaDB() {
	client := GetClient()
	angelinaDB := client.Database("angelina")

	err := angelinaDB.CreateCollection(context.TODO(), "settings")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
