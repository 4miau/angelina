package handlers

import (
	"angelina/events"

	"github.com/bwmarrin/discordgo"
)

func CaptureEvents(s *discordgo.Session) {
	s.AddHandler(events.Ready)
	s.AddHandler(events.MessageCreate)
}
