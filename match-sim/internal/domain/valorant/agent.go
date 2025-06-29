package valorant

func NewAgent(id int32, name AgentName, role AgentRole, abilities []*Ability) *Agent {
	return &Agent{
		Id:        id,
		Name:      name,
		Role:      role,
		Abilities: abilities,
	}
}

func CreateAgents() (agents []*Agent) {
	agents = append(agents, NewAgent(1, AgentName_ASTRA, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(2, AgentName_BREACH, AgentRole_DUELIST, []*Ability{}))
	agents = append(agents, NewAgent(3, AgentName_BRIMSTONE, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(4, AgentName_CHAMBER, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(5, AgentName_CLOVE, AgentRole_DUELIST, []*Ability{}))
	agents = append(agents, NewAgent(6, AgentName_HARBOR, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(7, AgentName_ISO, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(8, AgentName_JETT, AgentRole_DUELIST, []*Ability{}))
	agents = append(agents, NewAgent(9, AgentName_KAYO, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(10, AgentName_KILLJOY, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(11, AgentName_NEON, AgentRole_DUELIST, []*Ability{}))
	agents = append(agents, NewAgent(12, AgentName_RAZE, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(13, AgentName_REYNA, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(14, AgentName_SAGE, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(15, AgentName_SKYE, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(16, AgentName_TEJO, AgentRole_DUELIST, []*Ability{}))
	agents = append(agents, NewAgent(17, AgentName_VIPER, AgentRole_CONTROLLER, []*Ability{}))
	agents = append(agents, NewAgent(18, AgentName_WAYLAY, AgentRole_SENTINEL, []*Ability{}))
	agents = append(agents, NewAgent(19, AgentName_YORU, AgentRole_DUELIST, []*Ability{}))

	return agents
}
