package admin

import (
	"io"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

var SaveResult *int

// SaveServer executes the save command on the server.
func SaveServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	*SaveResult = 1
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Server received save command..")
	io.WriteString(*P, "/save\n")
	for {
		if *SaveResult == 2 {
			s.ChannelMessageSend(support.Config.FactorioChannelID, ".. saved successfully!")
			*SaveResult = 0
			break
		}
	}
	return
}
