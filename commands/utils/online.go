package utils

import (
	"container/list"
	"fmt"
	"io"
	"time"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// P references the var Pipe in main
var P *io.WriteCloser

var UserListResult *int
var UserList **list.List

func onlineListEmbed() *discordgo.MessageEmbed {
	*UserListResult = 1
	(**UserList).Init()

	io.WriteString(*P, "/p o\n")
	time.Sleep(600 * time.Millisecond)
	usercount := -1
	for {
		if *UserListResult == 3 {
			usercount = (**UserList).Len()
			*UserListResult = 0
			break
		} else if *UserListResult == 4 {
			usercount = 0
			*UserListResult = 0
			break
		}
	}

	fields := []*discordgo.MessageEmbedField{}
	var S = "player is"
	var S2 = " "
	if usercount != 1 {
		S = "players are"
		S2 = "s"
	}
	if usercount != 0 {
		for user := (**UserList).Front(); user != nil; user = user.Next() {
			var value string

			value = user.Value.(string)

			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  value,
				Value: "online",
			})
		}
	}
	embed := &discordgo.MessageEmbed{
		Type:        "rich",
		Color:       52,
		Description: fmt.Sprintf("%d %s online.", usercount, S),
		Title:       fmt.Sprintf("Online User%s", S2),
		Fields:      fields,
	}

	return embed

}

// ModsList returns the list of mods running on the server.
func OnlineList(s *discordgo.Session, m *discordgo.MessageCreate) {

	_, err := s.ChannelMessageSendEmbed(support.Config.FactorioChannelID, onlineListEmbed())
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID, fmt.Sprintf("Sorry, there was an error with the discord embed message Error details: %s", err))
	}
	return
}
