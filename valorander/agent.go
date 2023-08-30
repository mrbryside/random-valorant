package valorander

import (
	"errors"
	"math/rand"
	"strings"
)

func RandomAgent(role string) (string, error) {
	role = strings.ToLower(role)
	agent := NewAgent()
	mapping := map[string]PushPopAble{
		"controller": agent.controller,
		"sentinel":   agent.sentinel,
		"initiator":  agent.initiator,
		"duelist":    agent.duelist,
	}
	role_mapping := mapping[role]
	if role_mapping == nil {
		return "ไม่พบ Role: " + role, errors.New("error")
	}
	idx := rand.Intn(role_mapping.Len())

	result := role_mapping.GetByIdx(idx)
	return result, nil
}
