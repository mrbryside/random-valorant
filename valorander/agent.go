package valorander

import (
	"errors"
	"math/rand"
	"strings"
)

var (
	imageMapper = map[string]string{
		"astra":     "https://www.valorantpicker.com/assets/imgs/agents/portraits/astra.png",
		"brimstone": "https://www.valorantpicker.com/assets/imgs/agents/portraits/brimstone.png",
		"harbor":    "https://www.valorantpicker.com/assets/imgs/agents/portraits/harbor.png",
		"omen":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/omen.png",
		"viper":     "https://www.valorantpicker.com/assets/imgs/agents/portraits/viper.png",
		"chamber":   "https://www.valorantpicker.com/assets/imgs/agents/portraits/chamber.png",
		"cypher":    "https://www.valorantpicker.com/assets/imgs/agents/portraits/cypher.png",
		"deadlock":  "https://www.valorantpicker.com/assets/imgs/agents/portraits/deadlock.png",
		"killjoy":   "https://www.valorantpicker.com/assets/imgs/agents/portraits/killjoy.png",
		"sage":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/sage.png",
		"breach":    "https://www.valorantpicker.com/assets/imgs/agents/portraits/breach.png",
		"fade":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/fade.png",
		"gekko":     "https://www.valorantpicker.com/assets/imgs/agents/portraits/gekko.png",
		"kay/o":     "https://www.valorantpicker.com/assets/imgs/agents/portraits/kayo.png",
		"skye":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/skye.png",
		"sova":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/sova.png",
		"jett":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/jett.png",
		"neon":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/neon.png",
		"phoenix":   "https://www.valorantpicker.com/assets/imgs/agents/portraits/phoenix.png",
		"raze":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/raze.png",
		"reyna":     "https://www.valorantpicker.com/assets/imgs/agents/portraits/reyna.png",
		"yoru":      "https://www.valorantpicker.com/assets/imgs/agents/portraits/yoru.png",
	}
	imageIconMapper = map[string]string{
		"astra":     "https://www.valorantpicker.com/assets/imgs/agents/icons/astra.png",
		"brimstone": "https://www.valorantpicker.com/assets/imgs/agents/icons/brimstone.png",
		"harbor":    "https://www.valorantpicker.com/assets/imgs/agents/icons/harbor.png",
		"omen":      "https://www.valorantpicker.com/assets/imgs/agents/icons/omen.png",
		"viper":     "https://www.valorantpicker.com/assets/imgs/agents/icons/viper.png",
		"chamber":   "https://www.valorantpicker.com/assets/imgs/agents/icons/chamber.png",
		"cypher":    "https://www.valorantpicker.com/assets/imgs/agents/icons/cypher.png",
		"deadlock":  "https://www.valorantpicker.com/assets/imgs/agents/icons/deadlock.png",
		"killjoy":   "https://www.valorantpicker.com/assets/imgs/agents/icons/killjoy.png",
		"sage":      "https://www.valorantpicker.com/assets/imgs/agents/icons/sage.png",
		"breach":    "https://www.valorantpicker.com/assets/imgs/agents/icons/breach.png",
		"fade":      "https://www.valorantpicker.com/assets/imgs/agents/icons/fade.png",
		"gekko":     "https://www.valorantpicker.com/assets/imgs/agents/icons/gekko.png",
		"kay/o":     "https://www.valorantpicker.com/assets/imgs/agents/icons/kayo.png",
		"skye":      "https://www.valorantpicker.com/assets/imgs/agents/icons/skye.png",
		"sova":      "https://www.valorantpicker.com/assets/imgs/agents/icons/sova.png",
		"jett":      "https://www.valorantpicker.com/assets/imgs/agents/icons/jett.png",
		"neon":      "https://www.valorantpicker.com/assets/imgs/agents/icons/neon.png",
		"phoenix":   "https://www.valorantpicker.com/assets/imgs/agents/icons/phoenix.png",
		"raze":      "https://www.valorantpicker.com/assets/imgs/agents/icons/raze.png",
		"reyna":     "https://www.valorantpicker.com/assets/imgs/agents/icons/reyna.png",
		"yoru":      "https://www.valorantpicker.com/assets/imgs/agents/icons/yoru.png",
	}
)

func RandomAgent(role string) (string, string, string, error) {
	role = strings.ToLower(role)
	agent := NewAgent()
	mapping := map[string]PushPopAble{
		"controller": agent.Controller,
		"sentinel":   agent.Sentinel,
		"initiator":  agent.Initiator,
		"duelist":    agent.Duelist,
		"flex":       agent.Flex,
	}
	role_mapping := mapping[role]
	if role_mapping == nil {
		return "ไม่พบ Role: " + role, "", "", errors.New("error")
	}
	idx := rand.Intn(role_mapping.Len())

	result := role_mapping.GetByIdx(idx)
	return result, imageMapper[strings.ToLower(result)], imageIconMapper[strings.ToLower(result)], nil
}
