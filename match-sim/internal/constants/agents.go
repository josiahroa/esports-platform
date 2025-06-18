package constants

const (
	AgentRole_Controller = "Controller"
	AgentRole_Duelist    = "Duelist"
	AgentRole_Initiator  = "Initiator"
	AgentRole_Sentinel   = "Sentinel"
)

type Agent struct {
	ID        string
	Name      string
	Role      string
	Abilities []Ability
}

type Ability struct {
	Name            string
	CooldownSeconds int
	DurationSeconds int
	Damage          int
}

var Agents = []Agent{
	{
		ID:   "1",
		Name: "Astra",
		Role: AgentRole_Controller,
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
	{
		ID:   "2",
		Name: "Breach",
		Role: AgentRole_Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "3",
		Name: "Brimstone",
		Role: AgentRole_Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "4",
		Name: "Chamber",
		Role: AgentRole_Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{ID: "5",
		Name: "Clove",
		Role: AgentRole_Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "10",
		Name: "Harbor",
		Role: AgentRole_Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "11",
		Name: "Iso",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "12",
		Name: "Jett",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "13",
		Name: "KAY/O",
		Role: AgentRole_Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "14",
		Name: "Killjoy",
		Role: AgentRole_Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "15",
		Name: "Neon",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "18",
		Name: "Raze",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "19",
		Name: "Reyna",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "20",
		Name: "Sage",
		Role: AgentRole_Sentinel,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "21",
		Name: "Skye",
		Role: AgentRole_Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "23",
		Name: "Tejo",
		Role: AgentRole_Initiator,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "24",
		Name: "Viper",
		Role: AgentRole_Controller,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "25",
		Name: "Waylay",
		Role: AgentRole_Duelist,
		Abilities: []Ability{
			{
				Name:            "AbilityOne",
				CooldownSeconds: 1,
				DurationSeconds: 1,
				Damage:          1,
			},
		},
	},
	{
		ID:   "26",
		Name: "Yoru",
		Role: AgentRole_Duelist,
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
