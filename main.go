package main

import (
	"angelina/database"
	"angelina/handlers"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var (
	token *string
	dbUri string
)

func main() {
	err := godotenv.Load() // Loads dotenv file
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}

	envs, err := godotenv.Read()
	if err != nil {
		panic(err)
	}

	for key, val := range envs {
		if key == "token" {
			token = &val
		}
		if key == "dbServer" {
			dbUri = val
		}
	}

	s, err := discordgo.New(*token)
	if err != nil {
		fmt.Println("Err ", err.Error())
	}

	database.ConnectClient(dbUri)
	database.ConnectDB()
	handlers.CaptureEvents(s)

	wsError := s.Open()
	if wsError != nil {
		fmt.Println("Error opening WebSocket connection.")
		panic(wsError)
	}

	// Prevent bot from stopping after executing unless something is done to cancel it such as using Ctrl+C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	err = s.Close()
	if err != nil {
		panic(err)
	}
}
