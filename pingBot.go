package main

import (
	"flag"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func main() {

    BotID = u.ID
    
    dg.AddHandler(messageCreate)
    
    err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
  
  fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	// Simple way to keep program running until CTRL-C is pressed.
	<-make(chan struct{})
	return
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
		return
	}

	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
