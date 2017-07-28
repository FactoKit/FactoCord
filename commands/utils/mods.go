package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/FactoKit/FactoCord/support"
	"github.com/bwmarrin/discordgo"
)

type ModJson struct {
	Mods []Mod
}

type Mod struct {
	Name    string
	Enabled bool
}

func ModListEmbed(ModList *ModJson) *discordgo.MessageEmbed {
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

func ModsList(s *discordgo.Session, m *discordgo.MessageCreate) {
	ModList := &ModJson{}
	Json, err := ioutil.ReadFile(support.Config.ModListLocation)

	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID, "Error reading mod list, did you specify it in .env?")
		return
	}

	err = json.Unmarshal(Json, &ModList)
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID, fmt.Sprintf("%s", err))
		return
	}
	_, err = s.ChannelMessageSendEmbed(support.Config.FactorioChannelID, ModListEmbed(ModList))
	if err != nil {
		s.ChannelMessageSend(support.Config.FactorioChannelID, fmt.Sprintf("%s", err))
	}
	return
}
