package bot

import (
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
		cmd, _ := strings.CutPrefix(m.Content, config.Prefix)
		if cmd == "suck" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "hehe good boy")
		}
	}
}
