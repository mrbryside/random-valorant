package randomize

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"github.com/go-resty/resty/v2"
)

const MaxPlayer = 5

type PushPopAble interface {
	Pop(idx int) string
	Push(s string)
	GetByIdx(idx int) string
	GetAll() []string
	Len() int
}

type PushPop struct {
	data []string
}

func (p *PushPop) Pop(idx int) string {
	s := p.data[idx]
	p.data = append(p.data[:idx], p.data[idx+1:]...)
	return s
}

func (p *PushPop) Push(s string) {
	p.data = append(p.data, s)
}

func (p *PushPop) GetByIdx(idx int) string {
	return p.data[idx]
}

func (p *PushPop) GetAll() []string {
	return p.data
}

func (p *PushPop) Len() int {
	return len(p.data)
}

func (p *PushPop) Contains(s string) bool {
	for _, v := range p.data {
		if v == s {
			return true
		}
	}
	return false
}

type Controller PushPopAble
type Sentinel PushPopAble
type Initiator PushPopAble
type Duelist PushPopAble
type Player PushPopAble

type Agent struct {
	controller Controller
	sentinel   Sentinel
	initiator  Initiator
	duelist    Duelist
}

func NewAgent() Agent {
	offlineAgent := Agent{
		controller: &PushPop{data: []string{"Brimstone", "Harbor", "Omen", "Viper", "Astra"}},
		sentinel:   &PushPop{data: []string{"Chamber", "Cypher", "Deadlock", "Killjoy", "Sage"}},
		initiator:  &PushPop{data: []string{"Breach", "Fade", "Gekko", "Kay/o", "Skye", "Sova"}},
		duelist:    &PushPop{data: []string{"Jett", "Neon", "Phoenix", "Raze", "Reyna", "Yoru"}},
	}
	// fetch agent from api
	client := resty.New()
	resp, err := client.R().Get("https://valorant-api.com/v1/agents")
	if err != nil {
		log.Println("Error for fetch api:", err)
		return offlineAgent
	}

	var agents AgentResList
	if err := json.Unmarshal(resp.Body(), &agents); err != nil {
		log.Println("Error unmarshal from fetch api:", err)
		return offlineAgent
	}
	controller := make([]string, 0)
	sentinel := make([]string, 0)
	initiator := make([]string, 0)
	duelist := make([]string, 0)

	for _, agent := range agents.Data {
		if agent.Role.DisplayName == "Controller" {
			controller = append(controller, agent.DisplayName)
		} else if agent.Role.DisplayName == "Sentinel" {
			sentinel = append(sentinel, agent.DisplayName)
		} else if agent.Role.DisplayName == "Initiator" {
			initiator = append(initiator, agent.DisplayName)
		} else if agent.Role.DisplayName == "Duelist" {
			duelist = append(duelist, agent.DisplayName)
		}
	}
	return Agent{
		controller: &PushPop{data: controller},
		sentinel:   &PushPop{data: sentinel},
		initiator:  &PushPop{data: initiator},
		duelist:    &PushPop{data: duelist},
	}

}

func (a *Agent) RandomRole() PushPopAble {
	idx := rand.Intn(4)
	var result PushPopAble
	switch idx {
	case 0:
		result = a.controller
	case 1:
		result = a.sentinel
	case 2:
		result = a.initiator
	case 3:
		result = a.duelist
	}
	return result
}

func (a *Agent) GetRandomList() []PushPopAble {
	result := make([]PushPopAble, 0)
	result = append(result, a.controller)
	result = append(result, a.sentinel)
	result = append(result, a.initiator)
	result = append(result, a.duelist)
	result = append(result, a.RandomRole())
	return result
}

type Team struct {
	p1 string
	p2 string
	p3 string
	p4 string
	p5 string
}

func NewTeam(p1 string, p2 string, p3 string, p4 string, p5 string, p6 string, p7 string, p8 string, p9 string, p10 string) (Team, Team) {
	mapping := map[int]string{
		0: p1,
		1: p2,
		2: p3,
		3: p4,
		4: p5,
		5: p6,
		6: p7,
		7: p8,
		8: p9,
		9: p10,
	}
	player := &PushPop{data: []string{}}
	for {
		idx := rand.Intn(10)
		p := mapping[idx]
		if !player.Contains(p) {
			player.Push(p)
		}
		if player.Len() == 10 {
			break
		}
	}
	t1 := Team{
		p1: player.GetByIdx(0),
		p2: player.GetByIdx(1),
		p3: player.GetByIdx(2),
		p4: player.GetByIdx(3),
		p5: player.GetByIdx(4),
	}
	t2 := Team{
		p1: player.GetByIdx(5),
		p2: player.GetByIdx(6),
		p3: player.GetByIdx(7),
		p4: player.GetByIdx(8),
		p5: player.GetByIdx(9),
	}
	return t1, t2
}

type PlayerResult struct {
	PlayerName string
	AgentName  string
}
type RandomResult []PlayerResult

type Roulette struct {
	player     Player
	agent      Agent
	teamResult RandomResult
}

func (r *Roulette) RandomPlayer() Player {
	random := &PushPop{data: make([]string, 0)}
	for {
		idx := rand.Intn(r.player.Len())
		player := r.player.GetByIdx(idx)

		if !random.Contains(player) {
			random.Push(player)
		}
		if random.Len() == MaxPlayer {
			break
		}
	}
	r.player = random
	return r.player
}

func (r *Roulette) RandomAgent() {
	list := r.agent.GetRandomList()
	for idx, player := range r.player.GetAll() {
		agent := list[idx]
		rd := rand.Intn(agent.Len())
		result := agent.Pop(rd)
		pr := PlayerResult{PlayerName: player, AgentName: result}
		r.teamResult = append(r.teamResult, pr)
	}
}

func (r *Roulette) GetResult() RandomResult {
	return r.teamResult
}

func NewRoulette(t Team) *Roulette {
	return &Roulette{player: &PushPop{data: []string{t.p1, t.p2, t.p3, t.p4, t.p5}}, agent: NewAgent()}
}

type PlayerGroup struct {
	P1  string
	P2  string
	P3  string
	P4  string
	P5  string
	P6  string
	P7  string
	P8  string
	P9  string
	P10 string
}

func RandomFacade(pg PlayerGroup) (RandomResult, RandomResult) {
	t1, t2 := NewTeam(pg.P1, pg.P2, pg.P3, pg.P4, pg.P5, pg.P6, pg.P7, pg.P8, pg.P9, pg.P10)
	// team 1
	r1 := NewRoulette(t1)
	r1.RandomPlayer()
	r1.RandomAgent()
	// team 2
	r2 := NewRoulette(t2)
	r2.RandomPlayer()
	r2.RandomAgent()
	return r1.GetResult(), r2.GetResult()
}

func PrintDiscordResult(pg PlayerGroup) string {
	r1, r2 := RandomFacade(pg)
	var result string
	result += "ผลการสุ่ม Agent สำหรับทีมที่ 1 และ 2 ได้ตามนี้ครับ\n\n"
	result += "ทีมที่ 1\n"
	result += "------------------\n"
	for _, v := range r1 {
		result += fmt.Sprintf("%s -> %s\n", v.PlayerName, v.AgentName)
	}
	result += "\nทีมที่ 2\n"
	result += "------------------\n"
	for _, v := range r2 {
		result += fmt.Sprintf("%s -> %s\n", v.PlayerName, v.AgentName)
	}
	result += "------------------\n"
	return result
}