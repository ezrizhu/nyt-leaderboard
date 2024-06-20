package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

const (
	WGreen = iota
	WYellow
	WGray
)

type Wordle [][5]int

func wordleInit(msg string) Wordle {
	lines := strings.Split(msg, "\n")
	lines = lines[2:]

	w := make(Wordle, len(lines))

	for i, line := range lines {
		j := 0
		for _, char := range line {
			switch char {
			case '🟩':
				w[i][j] = WGreen
			case '🟨':
				w[i][j] = WYellow
			case '⬛':
				w[i][j] = WGray
			}
			j += 1
		}
	}
	return w
}

func (w Wordle) success() bool {
	last := w[len(w)-1]
	for _, i := range last {
		if i != WGreen {
			return false
		}
	}
	return true
}

func (w Wordle) scoring() int {
	if !w.success() {
		return 0
	}

	return 7 - len(w)
}

func wordle(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Content
	user := m.Author.ID

	w := wordleInit(msg)
	score := w.scoring()
	scoreStr := strconv.Itoa(score)

	newScore := wl.addScore(user, score)

	newScoreStr := strconv.Itoa(newScore)
	resp := "You earned " + scoreStr + " points today\n"
	resp += "Total score: " + newScoreStr

	dt := time.Now()
	log.Info().
		Str("UserID", user).
		Str("Username", m.Author.Username).
		Str("Date", dt.String()).
		Int("Score", score).
		Int("Total", newScore).
		Msg("New Wordle Entry")

	s.ChannelMessageSendReply(m.ChannelID, resp, m.Reference())

	return
}
