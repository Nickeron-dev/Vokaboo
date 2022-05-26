package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Flags
var (
	BotToken = flag.String("token", "OTc5Mzg4OTI4Mjg3MTk5MzIy.GY0nX_.cB6mtqoxEQ1nk9AUkjbeKh9qkgnR2xzqlupjmI", "Bot token")
)

const timeout time.Duration = time.Second * 10

var games map[string]time.Time = make(map[string]time.Time)

func init() { flag.Parse() }

func main() {
	s, _ := discordgo.New("Bot " + *BotToken)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.Contains(m.Content, "$start") {

      
			if ch, err := s.State.Channel(m.ChannelID); err != nil || !ch.IsThread() {
				thread, err := s.MessageThreadStartComplex(m.ChannelID, m.ID, &discordgo.ThreadStart{
					Name:                "Learn words with: " + m.Author.Username,
					Invitable:           false,
					RateLimitPerUser:    5,
				})
				if err != nil {
					panic(err)
				}
        s.ChannelMessageDelete(thread.ID, m.ID)
				_, _ = s.ChannelMessageSend(thread.ID, "OK, starting...!")
				m.ChannelID = thread.ID
        go func() {
          /*
          timer := time.NewTimer(10 * time.Second)
          <-timer.C */
          s.ChannelMessageSend(thread.ID, "1" + m.Author.Username)
          time.Sleep(2 * time.Second)
          s.ChannelMessageSend(thread.ID, "2" + m.Author.Username)
          time.Sleep(2 * time.Second)
          s.ChannelMessageSend(thread.ID, "3" + m.Author.Username)
        }()
			} else {
				_, _ = s.ChannelMessageSendReply(m.ChannelID, "correct!", m.Reference())
        s.ChannelMessageDelete(m.ChannelID, m.ID)
        go func() {
          /*
          timer := time.NewTimer(10 * time.Second)
          <-timer.C */
          
          s.ChannelMessageSend(m.ChannelID, "1" + m.Author.Username)
          time.Sleep(2 * time.Second)
          s.ChannelMessageSend(m.ChannelID, "2" + m.Author.Username)
          time.Sleep(2 * time.Second)
          s.ChannelMessageSend(m.ChannelID, "3" + m.Author.Username)
        }()
			}
			games[m.ChannelID] = time.Now()

      
		}
	})
	s.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged)

	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}
	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Graceful shutdown")

}
