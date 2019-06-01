package admin

import (
	"io"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// SaveServer executes the save command on the server.
func SaveServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	io.WriteString(*P, "/save\n")
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Server saved successfully!")
	return
}
