package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var wl WordleLeaderboard
var cl ConnectionsLeaderboard

func init() {
	wl = wordleLeaderboardInit()
	cl = connectionsLeaderboardInit()
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	token := os.Getenv("dtoken")
	if token == "" {
		panic("no token")
	}

	// Init bot
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Discord session")
		return
	}
	dg.Identify.Intents |= discordgo.IntentsGuildMessages

	dg.AddHandler(newMsg)

	err = dg.Open()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open connection")
		return
	}
	log.Info().Msg("Bot is running")

	// Handling SIGINT
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	log.Info().Msg("Shutting down")
	dg.Close()
}
