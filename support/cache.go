package support

import (
	"log"

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
		log.Fatal(err)
	}
	GuildID := GuildChannel.GuildID
	members, err := s.State.Guild(GuildID)
	if err != nil {
		log.Fatal(err)
	}
	for _, member := range members.Members {
		Users = append(Users, UserList{UserID: member.User.ID, Nick: member.Nick,
			User: member.User})
	}
}
