package support

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// SearchForUser searches for the user to be mentioned.
func SearchForUser(name string) *discordgo.User {
	name = strings.Replace(name, "@", "", -1)
	for _, user := range Users {
		if strings.ToLower(user.Nick) == strings.ToLower(name) ||
			strings.ToLower(user.User.Username) == strings.ToLower(name) {
			return user.User
		}
	}
	return nil
}

// LocateMentionPosition locates the position in a string list for the discord mention.
func LocateMentionPosition(List []string) []int {
	positionlist := []int{}
	for i, String := range List {
		if strings.Contains(String, "@") {
			positionlist = append(positionlist, i)
		}
	}
	return positionlist
}
