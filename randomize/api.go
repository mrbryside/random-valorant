package randomize

type Role struct {
	DisplayName string `json:"displayName"`
}

type AgentRes struct {
	DisplayName string `json:"displayName"`
	Role        Role   `json:"role"`
}

type AgentResList struct {
	Data []AgentRes `json:"data"`
}
