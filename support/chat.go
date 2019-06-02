package support

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
)

// Chat pipes in-game chat to Discord.
func Chat(s *discordgo.Session, _ *discordgo.MessageCreate) {
	idMap := map[int]string{}

	r_join := regexp.MustCompile("received stateChanged peerID\\((\\d+)\\) oldState\\(WaitingForCommandToStartSendingTickClosures\\) newState\\(InGame\\)")
	r_disconnect := regexp.MustCompile("nextHeartbeatSequenceNumber\\(\\d+\\) removing peer\\((\\d+)\\).")

	fmt.Println("함수 연결됨")

	newcomer := -1
	left := ""
	for {
		t, err := tail.TailFile("factorio.log", tail.Config{Follow: true})
		if err != nil {
			ErrorLog(fmt.Errorf("%s: An error occurred when attempting to tail factorio.log\nDetails: %s", time.Now(), err))
		}
		for line := range t.Lines {
			//fmt.Println("메시지 : " + line.Text)
			if strings.Contains(line.Text, "[CHAT]") || strings.Contains(line.Text, "[JOIN]") || strings.Contains(line.Text, "[LEAVE]") {
				if !strings.Contains(line.Text, "<server>") {
					if strings.Contains(line.Text, "[JOIN]") {
						TmpList := strings.Split(line.Text, " ")
						// Don't hard code the channelID! }:<
						fmt.Println("newcomer(saved) : ", newcomer)
						idMap[newcomer] = TmpList[3]
						fmt.Println("newcomer(saved) : ", newcomer)
						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", idMap[newcomer]+" joined the game."))
						newcomer = -1
					} else if strings.Contains(line.Text, "[LEAVE]") {
						TmpList := strings.Split(line.Text, " ")
						left = TmpList[3]
					} else {
						TmpList := strings.Split(line.Text, " ")
						TmpList[3] = strings.Replace(TmpList[3], ":", "", -1)
						if strings.Contains(strings.Join(TmpList, " "), "@") {
							index := LocateMentionPosition(TmpList)

							for _, position := range index {
								User := SearchForUser(TmpList[position])

								if User == nil {
									continue
								}
								TmpList[position] = User.Mention()
							}

						}
						s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("<%s>: %s", TmpList[3], strings.Join(TmpList[4:], " ")))
					}

					/*
						 if strings.Contains(line.Text, "[LEAVE]") {
							TmpList := strings.Split(line.Text, " ")
							// Don't hard code the channelID! }:<
							s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", TmpList[3] + "left the game."))
						} else {


						}

					*/
				}
			} else if strings.Contains(line.Text, "Matching server connection resumed") {

				s.ChannelMessageSend(Config.FactorioChannelID, "..Done. You can join now!")
			} else if r_join.MatchString(line.Text) {
				newcomer, _ = strconv.Atoi(r_join.FindStringSubmatch(line.Text)[1])
				fmt.Println("newcomer(to add) : ", newcomer)
				//idMap[newcomer] = ""
			} else if r_disconnect.MatchString(line.Text) {
				pid_left, _ := strconv.Atoi(r_disconnect.FindStringSubmatch(line.Text)[1])
				fmt.Println("pid_left : ", pid_left)
				if idMap[pid_left] == left {
					s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", idMap[pid_left]+" left the game."))
					left = ""
				} else {
					s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", idMap[pid_left]+" has disconnected from the game."))
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
