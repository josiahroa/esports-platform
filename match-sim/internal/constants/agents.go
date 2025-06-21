package constants

type AgentName uint8

const (
	Astra AgentName = iota
	Breach
	Brimstone
	Chamber
	Clove
	Harbor
	Iso
	Jett
	KAYO
	Killjoy
	Neon
	Raze
	Reyna
	Sage
	Skye
	Tejo
	Viper
	Waylay
	Yoru
)

var agentName = map[AgentName]string{
	Astra:     "Astra",
	Breach:    "Breach",
	Brimstone: "Brimstone",
	Chamber:   "Chamber",
	Clove:     "Clove",
	Harbor:    "Harbor",
	Iso:       "Iso",
	Jett:      "Jett",
	KAYO:      "KAY/O",
	Killjoy:   "Killjoy",
	Neon:      "Neon",
	Raze:      "Raze",
	Reyna:     "Reyna",
	Sage:      "Sage",
	Skye:      "Skye",
	Tejo:      "Tejo",
	Viper:     "Viper",
	Waylay:    "Waylay",
	Yoru:      "Yoru",
}

func (ss AgentName) String() string {
	return agentName[ss]
}

type AgentRole uint8

const (
	Controller AgentRole = iota
	Duelist
	Initiator
	Sentinel
)

var agentRole = map[AgentRole]string{
	Controller: "Controller",
	Duelist:    "Duelist",
	Initiator:  "Initiator",
	Sentinel:   "Sentinel",
}

func (ss AgentRole) String() string {
	return agentRole[ss]
}

type Agent struct {
	ID        string
	Name      AgentName
	Role      AgentRole
	Abilities []Ability
}

type Ability struct {
	Name            string
	CooldownSeconds int
	DurationSeconds int
	Damage          int
}

var Agents = map[AgentName]Agent{
	Astra: {
		ID:   "1",
		Name: Astra,
		Role: Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
			{
				Name:            "AbilityTwo",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Breach: {
		ID:   "2",
		Name: Breach,
		Role: Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Brimstone: {
		ID:   "3",
		Name: Brimstone,
		Role: Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Chamber: {
		ID:   "4",
		Name: Chamber,
		Role: Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Clove: {ID: "5",
		Name: Clove,
		Role: Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Harbor: {
		ID:   "10",
		Name: Harbor,
		Role: Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Iso: {
		ID:   "11",
		Name: Iso,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Jett: {
		ID:   "12",
		Name: Jett,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	KAYO: {
		ID:   "13",
		Name: KAYO,
		Role: Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Killjoy: {
		ID:   "14",
		Name: Killjoy,
		Role: Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Neon: {
		ID:   "15",
		Name: Neon,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Raze: {
		ID:   "18",
		Name: Raze,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Reyna: {
		ID:   "19",
		Name: Reyna,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Sage: {
		ID:   "20",
		Name: Sage,
		Role: Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Skye: {
		ID:   "21",
		Name: Skye,
		Role: Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Tejo: {
		ID:   "23",
		Name: Tejo,
		Role: Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Viper: {
		ID:   "24",
		Name: Viper,
		Role: Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Waylay: {
		ID:   "25",
		Name: Waylay,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	Yoru: {
		ID:   "26",
		Name: Yoru,
		Role: Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
}
