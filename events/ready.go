package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Bot is now online! Press CTRL+C to exit.")
	err := session.UpdateListeningStatus("Listening to doctor issuing new orders!")
	if err != nil {
	}
}
