package support

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
)

func Chat(s *discordgo.Session, m *discordgo.MessageCreate) {
	for true == true {
		t, err := tail.TailFile("factorio.log", tail.Config{Follow: true})
		if err != nil {
			log.Fatal(err)
		}
		for line := range t.Lines {
			if strings.Contains(line.Text, "[CHAT]") || strings.Contains(line.Text, "[JOIN]") || strings.Contains(line.Text, "[LEAVE]") {
				if !strings.Contains(line.Text, "[Discord]") {

					if strings.Contains(line.Text, "[JOIN]") ||
						strings.Contains(line.Text, "[LEAVE]") {
						TmpList := strings.Split(line.Text, " ")
						// Don't hard code the channelID! }:<
						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", strings.Join(TmpList[3:], " ")))
					} else {

						TmpList := strings.Split(line.Text, " ")
						TmpList[3] = strings.Replace(TmpList[3], ":", "", -1)
						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("<%s>: %s", TmpList[3], strings.Join(TmpList[4:], " ")))
					}
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
