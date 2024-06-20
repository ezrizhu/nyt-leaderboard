package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

const (
	CYellow = iota
	CGreen
	CBlue
	CPurple
)

type Connections [][4]int

func connectionsInit(msg string) Connections {
	lines := strings.Split(msg, "\n")
	lines = lines[2:]

	c := make(Connections, len(lines))

	for i, line := range lines {
		j := 0
		for _, char := range line {
			switch char {
			case '🟨':
				c[i][j] = CYellow
			case '🟩':
				c[i][j] = CGreen
			case '🟦':
				c[i][j] = CBlue
			case '🟪':
				c[i][j] = CPurple
			}
			j += 1
		}
	}
	return c
}

func (c Connections) scoring() int {
	success := 0
	for _, attempt := range c {
		fst := attempt[0]
		same := true

		for _, num := range attempt {
			if num != fst {
				same = false
				break
			}
		}
		if same {
			success += 1
		}
	}

	if success == 4 {
		return 8 - len(c) + success
	} else {
		return success
	}
}

func connections(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := m.Content
	user := m.Author.ID

	c := connectionsInit(msg)
	score := c.scoring()

	scoreStr := strconv.Itoa(score)

	newScore := cl.addScore(user, score)

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
		Msg("New Connections Entry")

	s.ChannelMessageSendReply(m.ChannelID, resp, m.Reference())

	return

}
