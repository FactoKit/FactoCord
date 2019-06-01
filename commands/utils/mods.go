package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"../../support"
	"github.com/bwmarrin/discordgo"
)

// ModJson is struct containing a slice of Mod.
type ModJson struct {
	Mods []Mod
}

// Mod is a struct containing info about a mod.
type Mod struct {
	Name    string
	Enabled bool
}

func modListEmbed(ModList *ModJson) *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{}
	var enabled, disabled int
	var S = "mod"
	if len(ModList.Mods) > 1 {
		S = "mods"
	}
	for _, mod := range ModList.Mods {
		var value string

		if mod.Enabled {
			value = "Enabled"
			enabled = enabled + 1
		} else {
			value = "Disabled"
			disabled = disabled + 1
		}

		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  mod.Name,
			Value: value,
		})
	}
	embed := &discordgo.MessageEmbed{
		Type:  "rich",
		Color: 52,
		Description: fmt.Sprintf("%d total %s (%d enabled, %d disabled)", len(ModList.Mods),
			S, enabled, disabled),
		Title:  "Mods",
		Fields: fields,
	}

	return embed

}

// ModsList returns the list of mods running on the server.
func ModsList(s *discordgo.Session, m *discordgo.MessageCreate) {
	ModList := &ModJson{}
	Json, err := ioutil.ReadFile(support.Config.ModListLocation)
	// Don't exit on this error, just sent message to the channel!
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID,
			fmt.Sprintf("Sorry, there was an error reading your mods list, did you specify it in the .env file? Error details: %s", err))
		return
	}

	err = json.Unmarshal(Json, &ModList)
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID,
			fmt.Sprintf("Sorry, there was an error reading your mods list. Error details: %s", err))
		return
	}
	_, err = s.ChannelMessageSendEmbed(support.Config.FactorioChannelID, modListEmbed(ModList))
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID, fmt.Sprintf("Sorry, there was an error with the discord embed message Error details: %s", err))
	}
	return
}
