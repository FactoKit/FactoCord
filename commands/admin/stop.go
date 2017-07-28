package admin

import (
	"io"
	"os"

	"github.com/FactoKit/FactorioCord/support"
	"github.com/bwmarrin/discordgo"
)

var P *io.WriteCloser

// Saves and stops the server
func StopServer(s *discordgo.Session, m *discordgo.MessageCreate) {
	io.WriteString(*P, "/save\n")
	io.WriteString(*P, "/quit\n")
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Server saved and shutting down; Cya!")
	s.Close()
	os.Exit(1)
}
