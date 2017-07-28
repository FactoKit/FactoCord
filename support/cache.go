package support

import (
	"errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type UserList struct {
	UserID string
	Nick   string
	User   *discordgo.User
}

var Users []UserList

// Caches the users list to be searched
func CacheDiscordMembers(s *discordgo.Session) {
	// Clear the users list
	Users = nil

	GuildChannel, err := s.Channel(Config.FactorioChannelID)
	if err != nil {
		ErrorLog(errors.New(fmt.Sprintf("%s: An error occurred when attempting to read the Discord Guild\nDetails: %s", time.Now(), err)))
	}
	GuildID := GuildChannel.GuildID
	members, err := s.State.Guild(GuildID)
	if err != nil {
		ErrorLog(errors.New(fmt.Sprintf("%s: An error occurred when attempting to read the Discord Guild Members\nDetails: %s", time.Now(), err)))
	}
	for _, member := range members.Members {
		Users = append(Users, UserList{UserID: member.User.ID, Nick: member.Nick,
			User: member.User})
	}
}
