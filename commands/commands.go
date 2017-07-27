package commands

import (
	"strings"

	"github.com/FM1337/FactorioCord/commands/admin"
	"github.com/FM1337/FactorioCord/commands/utils"
	"github.com/FM1337/FactorioCord/support"
	"github.com/bwmarrin/discordgo"
)

type Commands struct {
	CommandList []Command
}

type Command struct {
	Name    string
	Command func(s *discordgo.Session, m *discordgo.MessageCreate)
	Admin   bool
}

var CL Commands

// register commands on startup
func RegisterCommands() {
	CL.CommandList = append(CL.CommandList, Command{Name: "Stop", Command: admin.StopServer,
		Admin: true})
	CL.CommandList = append(CL.CommandList, Command{Name: "Mods", Command: utils.ModsList,
		Admin: false})
}

// run a command
func RunCommand(name string, s *discordgo.Session, m *discordgo.MessageCreate) {
	for _, command := range CL.CommandList {
		if strings.ToLower(command.Name) == strings.ToLower(name) {
			if command.Admin && CheckAdmin(m.Author.ID) {
				command.Command(s, m)
			}

			if !command.Admin {
				command.Command(s, m)
			}

			return
		}
	}
}

// Check if the user attempting to run an admin command is an admin
func CheckAdmin(ID string) bool {
	for _, admin := range support.Config.AdminIDs {
		if ID == admin {
			return true
		}
	}
	return false
}
