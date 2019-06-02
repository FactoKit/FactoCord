package support

import (
	"container/list"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/hpcloud/tail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//SaveResult : 0 : none, 1 : request save, 2 : save completed
var SaveResult int

var UserListResult int
var OnlineUserList *list.List

// Chat pipes in-game chat to Discord.
func Chat(s *discordgo.Session, _ *discordgo.MessageCreate) {
	SaveResult = 0
	UserListResult = 0

	OnlineUserList = list.New()

	idMap := map[int]string{}

	r_join := regexp.MustCompile("received stateChanged peerID\\((\\d+)\\) oldState\\(WaitingForCommandToStartSendingTickClosures\\) newState\\(InGame\\)")
	r_disconnect := regexp.MustCompile("nextHeartbeatSequenceNumber\\(\\d+\\) removing peer\\((\\d+)\\).")
	r_save := regexp.MustCompile("Info AppManagerStates.cpp:\\d+: Saving finished")
	r_online := regexp.MustCompile("Online players? \\((\\d+)\\)")
	onlineUserCount := 0
	r_online_name := regexp.MustCompile("(\\w+) \\(online\\)")

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
						idMap[newcomer] = TmpList[3]
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
				//idMap[newcomer] = ""
			} else if r_disconnect.MatchString(line.Text) {
				pid_left, _ := strconv.Atoi(r_disconnect.FindStringSubmatch(line.Text)[1])
				if idMap[pid_left] == left {
					s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", idMap[pid_left]+" left the game."))
					left = ""
				} else {
					s.ChannelMessageSend(Config.FactorioChannelID, fmt.Sprintf("%s", idMap[pid_left]+" has disconnected from the game."))
				}
			} else if SaveResult == 1 && r_save.MatchString(line.Text) {
				SaveResult = 2
			} else if UserListResult == 1 && r_online.MatchString(line.Text) {
				fmt.Println("유저리스트 요청 접수")
				UserListResult = 2
				onlineUserCount, _ = strconv.Atoi(r_online.FindStringSubmatch(line.Text)[1])
				if onlineUserCount == 0 {
					OnlineUserList.PushBack("No one is online..")
					UserListResult = 4
				}

			} else if UserListResult == 2 && r_online_name.MatchString(line.Text) {
				fmt.Println("유저리스트 요청 처리중, 유저 이름 나옴")
				onlineUserCount--
				OnlineUserList.PushBack(r_online_name.FindStringSubmatch(line.Text)[1])
				if onlineUserCount == 0 {
					fmt.Println("유저리스트 요청 리턴(있는경우)")
					UserListResult = 3
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
}
