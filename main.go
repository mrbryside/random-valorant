package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/mrbryside/valorant-random/valorander"
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

	goBot.AddHandler(messageTeamHandler)
	goBot.AddHandler(messageAgentHandler)
	goBot.AddHandler(messageMapHandler)

	return goBot
}

func messageTeamHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!team") {
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
			res += fmt.Sprintf("โปรดใส่ชื่อผู้เล่นให้พอดี 10 คน, คนละ 1 บรรทัดด้วยครับ \n(ตอนนี้มี %d คน)\n", len(msg))
			res += fmt.Sprintln("")
			res += fmt.Sprintln("**ตัวอย่าง")
			res += fmt.Sprintln("!team")
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
			res += fmt.Sprintln("")
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title:       "⛔  ข้อผิดพลาด",
					Description: res,
					Color:       0xcc241d,
				},
			})
			return
		}

		result := valorander.GenerateTeamResult(valorander.PlayerGroup{
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
		_, _ = s.ChannelMessageSendComplex(m.ChannelID, result)
	}
}

func messageAgentHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!agent") {
		msg := make([]string, 0)
		lines := strings.Split(m.Content, " ")
		for idx, line := range lines {
			if idx == 0 {
				continue
			}
			// Trim any leading/trailing spaces
			line = strings.TrimSpace(line)
			msg = append(msg, line)
		}
		if len(msg) != 1 {
			res := fmt.Sprintln("โปรดใส่ Role ที่ต้องการให้ Random ด้วยครับ")
			res += fmt.Sprintln("")
			res += fmt.Sprintln("**ตัวอย่าง")
			res += fmt.Sprintln("!agent controller")

			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title:       "⛔  ข้อผิดพลาด",
					Description: res,
					Color:       0xcc241d,
				},
			})
			return
		}

		agent, image, imageIcon, err := valorander.RandomAgent(msg[0])
		if err != nil {
			s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title:       "⛔  ข้อผิดพลาด",
					Description: agent,
					Color:       0xcc241d,
					Image: &discordgo.MessageEmbedImage{
						URL: image,
					},
				},
			})
			return
		}
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "📌 ผลการสุ่ม Agent",
				Description: agent,
				Color:       0x83a598,
				// Image: &discordgo.MessageEmbedImage{
				// 	URL: image,
				// },
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: imageIcon,
				},
			},
		})
	}
}

func messageMapHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.HasPrefix(m.Content, "!map") {

		result := valorander.RandomMap()
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "📌 ผลการสุ่ม Map",
				Description: result.Name,
				Color:       0x83a598,
				Image: &discordgo.MessageEmbedImage{
					URL: result.ImageUrl,
				},
			},
		})
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
