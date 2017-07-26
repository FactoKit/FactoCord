package support

import (
	"fmt"
	"os"
	"strings"
)

var Config config

type config struct {
	DiscordToken       string
	FactorioChannelID  string
	SaveFile           string
	ServerSettingsFile string
	Executable         string
	AdminIDs           []string
	Prefix             string
}

func (conf *config) LoadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("Enviroment file not found, cannot continue!")
		os.Exit(1)
	}

	Config = config{
		DiscordToken:       os.Getenv("DiscordToken"),
		FactorioChannelID:  os.Getenv("FactorioChannelID"),
		SaveFile:           os.Getenv("SaveFile"),
		ServerSettingsFile: os.Getenv("ServerSettings"),
		Executable:         os.Getenv("Executable"),
		AdminIDs:           strings.Split(os.Getenv("AdminIDs"), ","),
		Prefix:             os.Getenv("Prefix"),
	}

}
