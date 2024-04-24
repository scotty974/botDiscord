package bot

import (
	"fmt"

	"github.com/SteakBarbare/GoBot-Cours/config"
	"github.com/bwmarrin/discordgo"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.TOKEN)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.Content == config.BOT_PREFIX+"quoi" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "feur")
	}
	if m.Author.ID == "435053014093529099" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Il est le créateur !")
	}
}
