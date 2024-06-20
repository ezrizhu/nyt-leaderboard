package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func newMsg(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID != "1251996771912126525" {
		return
	}

	switch {
	case strings.HasPrefix(m.Content, "Wordle"):
		wordle(s, m)
	case strings.HasPrefix(m.Content, "Connections"):
		connections(s, m)
	case strings.HasPrefix(m.Content, "!wl"):
		wLeaderboard(s, m)
	case strings.HasPrefix(m.Content, "!cl"):
		cLeaderboard(s, m)
	default:
		return
	}
}

func wLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Wordle",
		Description: wl.showLeaderboard(),
		Color:       2227217,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func cLeaderboard(s *discordgo.Session, m *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       "Connections",
		Description: cl.showLeaderboard(),
		Color:       2227217,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
