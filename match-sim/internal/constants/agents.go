package constants

type AgentRole string

type Agent struct {
	ID                     string
	Name                   string
	Role                   AgentRole
	AbilityOne             string
	AbilityTwo             string
	AbilityThree           string
	UltimateAbility        string
	UltimateAbilityAltFire string
}

var Agents = []Agent{
	{ID: "1", Name: "Astra", Role: "Controller"},
	{ID: "2", Name: "Breach", Role: "Initiator"},
	{ID: "3", Name: "Brimstone", Role: "Controller"},
	{ID: "4", Name: "Chamber", Role: "Sentinel"},
	{ID: "5", Name: "Clove", Role: "Controller"},
	{ID: "6", Name: "Cypher", Role: "Sentinel"},
	{ID: "7", Name: "Deadlock", Role: "Sentinel"},
	{ID: "8", Name: "Fade", Role: "Initiator"},
	{ID: "9", Name: "Gekko", Role: "Initiator"},
	{ID: "10", Name: "Harbor", Role: "Controller"},
	{ID: "11", Name: "Iso", Role: "Duelist"},
	{ID: "12", Name: "Jett", Role: "Duelist"},
	{ID: "13", Name: "KAY/O", Role: "Initiator"},
	{ID: "14", Name: "Killjoy", Role: "Sentinel"},
	{ID: "15", Name: "Neon", Role: "Duelist"},
	{ID: "16", Name: "Omen", Role: "Controller"},
	{ID: "17", Name: "Phoenix", Role: "Duelist"},
	{ID: "18", Name: "Raze", Role: "Duelist"},
	{ID: "19", Name: "Reyna", Role: "Duelist"},
	{ID: "20", Name: "Sage", Role: "Sentinel"},
	{ID: "21", Name: "Skye", Role: "Initiator"},
	{ID: "22", Name: "Sova", Role: "Initiator"},
	{ID: "23", Name: "Tejo", Role: "Initiator"},
	{ID: "24", Name: "Viper", Role: "Controller"},
	{ID: "25", Name: "Waylay", Role: "Duelist"},
	{ID: "26", Name: "Yoru", Role: "Duelist"},
}
