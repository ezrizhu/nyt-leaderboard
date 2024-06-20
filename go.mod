module github.com/ezrizhu/roll

go 1.22

require github.com/rs/zerolog v1.32.0

require (
	github.com/bwmarrin/discordgo v0.27.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
	golang.org/x/sys v0.12.0 // indirect
)

require github.com/ezrizhu/roll/discordgo v0.0.0

replace github.com/ezrizhu/roll/discordgo => ./discordgo
