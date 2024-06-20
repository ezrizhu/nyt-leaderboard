package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func newMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID != "1253428089409503385" {
		return
	}

	switch {
	case strings.HasPrefix(m.Content, "Wordle"):
		wordle(s, m)
	case strings.HasPrefix(m.Content, "Connections"):
		connections(s, m)
	case strings.HasPrefix(m.Content, "!leaderboard"):
		//leaderboard(s, m)
		return
	default:
		return
	}
}
