package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
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
	fmt.Println(c)
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
	_ = user

	c := connectionsInit(msg)
	score := c.scoring()

	scoreStr := strconv.Itoa(score)

	s.ChannelMessageSendReply(m.ChannelID, scoreStr, m.Reference())

	// write to leaderboard
	return
}
