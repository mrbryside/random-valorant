package randomize

import "math/rand"

func RandomAgent(role string) string {
	agent := NewAgent()
	mapping := map[string]PushPopAble{
		"controller": agent.controller,
		"sentinel":   agent.sentinel,
		"initiator":  agent.initiator,
		"duelist":    agent.duelist,
	}
	role_mapping := mapping[role]
	idx := rand.Intn(role_mapping.Len())

	return role_mapping.GetByIdx(idx)
}
