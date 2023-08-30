package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mrbryside/valorant-random/randomize"
)

var (
	Token     string
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

var BotId string

func ReadConfig() {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
}

func Start() *discordgo.Session {
	goBot, err := discordgo.New("Bot " + Token)
	if err != nil {
		panic(err)
	}

	u, err := goBot.User("@me")
	if err != nil {
		panic(err)
	}
	BotId = u.ID

	goBot.AddHandler(messageHandler)

	return goBot
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!random") {
		msg := make([]string, 0)
		lines := strings.Split(m.Content, "\n")
		for idx, line := range lines {
			if idx == 0 {
				continue
			}
			// Trim any leading/trailing spaces
			line = strings.TrimSpace(line)
			msg = append(msg, line)
		}
		if len(msg) != 10 {
			res := ""
			res += fmt.Sprintf("[ข้อผิดพลาด] โปรดใส่ชื่อผู้เล่นให้พอดี 10 คน, คนละ 1 บรรทัดด้วยครับ ตอนนี้มี %d คน\n", len(msg))
			res += fmt.Sprintln("----------------------")
			res += fmt.Sprintln("**ตัวอย่าง")
			res += fmt.Sprintln("----------------------")
			res += fmt.Sprintln("!random")
			res += fmt.Sprintln("Bank")
			res += fmt.Sprintln("Guy")
			res += fmt.Sprintln("Jet")
			res += fmt.Sprintln("Care")
			res += fmt.Sprintln("Cherry")
			res += fmt.Sprintln("Ball")
			res += fmt.Sprintln("Kie")
			res += fmt.Sprintln("Ford")
			res += fmt.Sprintln("Add")
			res += fmt.Sprintln("Pi")
			res += fmt.Sprintln("----------------------")

			s.ChannelMessageSend(m.ChannelID, res)
			return
		}

		result := randomize.PrintDiscordResult(randomize.PlayerGroup{
			P1:  msg[0],
			P2:  msg[1],
			P3:  msg[2],
			P4:  msg[3],
			P5:  msg[4],
			P6:  msg[5],
			P7:  msg[6],
			P8:  msg[7],
			P9:  msg[8],
			P10: msg[9],
		})
		_, _ = s.ChannelMessageSend(m.ChannelID, result)
	}
}

func main() {
	ReadConfig()
	goBot := Start()
	err := goBot.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println("Bot is running")
	<-make(chan struct{})
	goBot.Close()
}
