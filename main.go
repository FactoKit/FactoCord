package main

import (
	"FactorioCord/commands"
	"FactorioCord/commands/admin"
	"FactorioCord/support"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var Running bool
var Pipe io.WriteCloser
var Session *discordgo.Session

func main() {
	support.Config.LoadEnv()
	Running = false
	err := os.Remove("factorio.log")

	if err != nil {
		log.Println(err)
	}

	logging, err := os.OpenFile("factorio.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal(err)
	}

	mwriter := io.MultiWriter(logging, os.Stdout)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		cmd := exec.Command(support.Config.Executable, support.Config.LaunchParameters...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = mwriter
		Pipe, err = cmd.StdinPipe()

		if err != nil {
			log.Fatal(err)
		}

		err := cmd.Start()

		if err != nil {
			log.Fatal(err)
		}

	}()

	go func() {
		Console := bufio.NewReader(os.Stdin)
		for true == true {
			line, _, err := Console.ReadLine()
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(Pipe, fmt.Sprintf("%s\n", line))
		}
	}()

	go func() {
		// Wait 10 seconds on start up before continuing
		time.Sleep(10 * time.Second)

		for true == true {
			support.CacheDiscordMembers(Session)
			//sleep for 4 hours (caches every 4 hours)
			time.Sleep(4 * time.Hour)
		}
	}()
	Discord()
}

func Discord() {
	// No hard coding the token }:<
	discordToken := support.Config.DiscordToken
	commands.RegisterCommands()
	admin.P = &Pipe
	fmt.Println("Starting bot..")
	bot, err := discordgo.New("Bot " + discordToken)
	Session = bot
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	err = bot.Open()

	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	bot.AddHandler(messageCreate)
	bot.AddHandlerOnce(support.Chat)
	time.Sleep(3 * time.Second)
	bot.ChannelMessageSend(support.Config.FactorioChannelID, "The server has started!")
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	log.Print("[" + m.Author.Username + "] " + m.Content)

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == support.Config.FactorioChannelID {
		if strings.HasPrefix(m.Content, support.Config.Prefix) {
			command := strings.Split(m.Content[1:len(m.Content)], " ")
			name := strings.ToLower(command[0])
			commands.RunCommand(name, s, m)
			return
		} else {
			// Pipes normal chat allowing it to be seen ingame
			io.WriteString(Pipe, fmt.Sprintf("[Discord] <%s>: %s\r\n", m.Author.Username, m.ContentWithMentionsReplaced()))
			return
		}
	}
}
