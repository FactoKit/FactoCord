package commands

import (
	"strings"

	"../support"
	"./admin"
	"./utils"
	"github.com/bwmarrin/discordgo"
)

// Commands is a struct containing a slice of Command.
type Commands struct {
	CommandList []Command
}

// Command is a struct containing fields that hold command information.
type Command struct {
	Name    string
	Command func(s *discordgo.Session, m *discordgo.MessageCreate)
	Admin   bool
}

// CL is a Commands interface.
var CL Commands

// RegisterCommands registers the commands on start up.
func RegisterCommands() {
	// Admin Commands
	CL.CommandList = append(CL.CommandList, Command{Name: "Stop", Command: admin.StopServer,
		Admin: true})
	CL.CommandList = append(CL.CommandList, Command{Name: "Restart", Command: admin.Restart,
		Admin: true})
	CL.CommandList = append(CL.CommandList, Command{Name: "Save", Command: admin.SaveServer,
		Admin: true})

	// Util Commands
	CL.CommandList = append(CL.CommandList, Command{Name: "Mods", Command: utils.ModsList,
		Admin: false})
	CL.CommandList = append(CL.CommandList, Command{Name: "Online", Command: utils.OnlineList,
		Admin: false})
}

// RunCommand runs a specified command.
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

// CheckAdmin checks if the user attempting to run an admin command is an admin
func CheckAdmin(ID string) bool {
	for _, admin := range support.Config.AdminIDs {
		if ID == admin {
			return true
		}
	}
	return false
}
