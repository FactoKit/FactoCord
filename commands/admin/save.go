package admin

import (
	"io"

	"github.com/FactoKit/FactoCord/support"
	"github.com/bwmarrin/discordgo"
)

// Saves the server
func SaveServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	io.WriteString(*P, "/save\n")
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Server saved successfully!")
	return
}
