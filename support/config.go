package support

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// Config is a config interface.
var Config config

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

type config struct {
	DiscordToken          string
	FactorioChannelID     string
	Executable            string
	LaunchParameters      []string
	AdminIDs              []string
	Prefix                string
	ModListLocation       string
	GameName              string
	PassConsoleChat       bool
	EnableConsoleChannel  bool
	FactorioConsoleChatID string
}

func getenvStr(key string) (string, error) {
    v := os.Getenv(key)
    if v == "" {
        return v, ErrEnvVarEmpty
    }
    return v, nil
}

func getenvBool(key string) (bool) {
    s, err := getenvStr(key)
    if err != nil {
        return false
    }
    v, err := strconv.ParseBool(s)
    if err != nil {
        return false
    }
    return v
}

func (conf *config) LoadEnv() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("Environment file not found, cannot continue!")
		Error := errors.New("Failed to load environment file")
		ErrorLog(Error)
	}
	Config = config{
		DiscordToken:            os.Getenv("DiscordToken"),
		FactorioChannelID:       os.Getenv("FactorioChannelID"),
		LaunchParameters:        strings.Split(os.Getenv("LaunchParameters"), " "),
		Executable:              os.Getenv("Executable"),
		AdminIDs:                strings.Split(os.Getenv("AdminIDs"), ","),
		Prefix:                  os.Getenv("Prefix"),
		ModListLocation:         os.Getenv("ModListLocation"),
		GameName:                os.Getenv("GameName"),
		PassConsoleChat:         getenvBool("PassConsoleChat"),
		EnableConsoleChannel:    getenvBool("EnableConsoleChannel"),
		FactorioConsoleChatID:   os.Getenv("FactorioConsoleChatID"),
	}
}
