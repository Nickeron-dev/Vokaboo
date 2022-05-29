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
  word_set := make(map[string]string)
  word_set["hallo"] = "hello"
  word_set["warum"] = "why"
  word_set["gut"] = "good"
	s, _ := discordgo.New("Bot " + *BotToken)
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		fmt.Println("Bot is ready")
	})
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if strings.Contains(m.Content, "$start") {
				_, _ = s.ChannelMessageSendReply(m.ChannelID, "correct!", m.Reference())
        s.ChannelMessageDelete(m.ChannelID, m.ID)
        go func() {
          /*
          timer := time.NewTimer(10 * time.Second)
          <-timer.C */
          t1 := time.NewTimer(time.Second * 15)
          quit := make(chan bool)
          go func() {
            s.ChannelMessageSend(m.ChannelID, "Timer run...")
            <- t1.C
            quit <- true
            s.ChannelMessageSend(m.ChannelID, "Time is up!")
          }()
          for {
              select {
              case <- quit:
                  return
              default:
                for key := range word_set {
                  s.ChannelMessageSend(m.ChannelID, "1" + word_set[key] + m.Author.Username)
                  s.AddHandlerOnce(func(s *discordgo.Session, m *discordgo.MessageCreate) {
                    fmt.Println("here")
                  })
                  time.Sleep(4 * time.Second)
                }
              }
          }
          
        }()
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
