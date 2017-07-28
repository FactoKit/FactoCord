package admin

import (
	"io"
	"time"

	"github.com/FM1337/FactorioCord/support"
	"github.com/bwmarrin/discordgo"
)

var R *bool
var RestartCount int

// Saves and restarts the server
func Restart(s *discordgo.Session, m *discordgo.MessageCreate) {
	if *R == false {
		s.ChannelMessageSend(support.Config.FactorioChannelID, "Server is not running!")
		return
	}
	io.WriteString(*P, "/save\n")
	io.WriteString(*P, "/quit\n")
	s.ChannelMessageSend(support.Config.FactorioChannelID, "Saved server, now restarting!")
	time.Sleep(3 * time.Second)
	*R = false
	RestartCount = RestartCount + 1
	return
}
