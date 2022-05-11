package bot

import (
	"errors"
	"fmt"
	"strings"
	"github.com/bwmarrin/discordgo"
	"github.com/kevin/boibot/config"
)

type discErrors struct {
	msg string
}

var users map[string]bool 
var outChannelID string;

var BotId string
var goBot *discordgo.Session

func Start() {
	users = make(map[string]bool);
	users["415213790867488790"] = true
	outChannelID = "876559963340804176"

	goBot, err := discordgo.New("Bot " + config.Token)

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
	fmt.Println("Bot is running !")
}
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}
	if m.ChannelID == "894255837147701318" || m.ChannelID == "635283165971480597" {
		if _, ok := users[m.Author.ID]; !ok {
		
			s.ChannelMessageDelete(m.ChannelID, m.ID);	
			if m.Author.ID == "71670326085292032" || m.Author.ID == "178205523114590208" {
				channel, _ := s.UserChannelCreate(m.Author.ID);
				s.ChannelMessageSend(channel.ID, fmt.Sprintf("Glorious Leader has seen you.\n"));
			} else {
			s.ChannelMessageSend(outChannelID, fmt.Sprintf("-15 Social Credit to %s for posting in Game Anouncements!\n", m.Author.Username));
		}
		}
	}
if m.ChannelID == "893613898673049600" {
		if _, ok := users[m.Author.ID]; !ok {
		
			s.ChannelMessageDelete(m.ChannelID, m.ID);	
			s.ChannelMessageSend(outChannelID, fmt.Sprintf("%s Has posted in Important Stuff!, -1,000,000 social credit! Execution date is tomorrow!\n", m.Author.Username));
		}
	}
	
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) (bool) {
	message := strings.Split(m.Content, " ");
	if message[0][0:1] == "!" {
		command := message[0][1:];
		switch command {
		case "AddUsers":
			stat, err := AddUser(m.Mentions);
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%s", err));
			} else {
				s.ChannelMessageSend(m.ChannelID, stat);
			}	
			return true	
		default:
			return false;
		}

	}
	return false;
}

func AddUser(uArr []*discordgo.User) (string, error) {
	if len(uArr) < 1 {
		return "", errors.New("No Users Supplied");
	}
	for _, u := range uArr {
		users[u.ID] = true;

	}
	return "Users succesfully added", nil;
}
