package support

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Config is a config interface.
var Config config

type config struct {
	DiscordToken      string
	FactorioChannelID string
	Executable        string
	LaunchParameters  []string
	AdminIDs          []string
	Prefix            string
	ModListLocation   string
	GameName          string
}

func (conf *config) LoadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("Environment file not found, cannot continue!")
		Error := errors.New("Failed to load environment file")
		ErrorLog(Error)
	}

	Config = config{
		DiscordToken:      os.Getenv("DiscordToken"),
		FactorioChannelID: os.Getenv("FactorioChannelID"),
		LaunchParameters:  strings.Split(os.Getenv("LaunchParameters"), " "),
		Executable:        os.Getenv("Executable"),
		AdminIDs:          strings.Split(os.Getenv("AdminIDs"), ","),
		Prefix:            os.Getenv("Prefix"),
		ModListLocation:   os.Getenv("ModListLocation"),
		GameName:          os.Getenv("GameName"),
	}

}
