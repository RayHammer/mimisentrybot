package bot

import (
	"container/list"
	"fmt"
	"log"
	"mimisentry/config"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
var cmds list.List

func Run() bool {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return false
	}

	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return false
	}
	BotID = user.ID
	goBot.AddHandler(messageCreate)

	err = goBot.Open()
	if err != nil {
		return false
	}

	cmds = *list.New()
	cmds.PushBack("")

	fmt.Println("The abomination is running. Press Ctrl-C to kill it.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	goBot.Close()

	return true
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	if strings.HasPrefix(m.Content, config.Prefix) {
		cmd := strings.TrimPrefix(m.Content, config.Prefix)
		switch {
		case strings.HasPrefix(cmd, "help"):
			_, _ = s.ChannelMessageSend(m.ChannelID, "no one's gonna come to help you")
		case strings.HasPrefix(cmd, "suck"):
			_, _ = s.ChannelMessageSend(m.ChannelID, "hehe good boy")
			break
		}
	}
}
