package main

var AllPaths = map[string]Path{
	"9": {
		Name:        "9",
		Concept:     "",
		Connections: "FBI Agent, Safe House Owner, Weapons Dealer, Lab Worker",
		Skills: [][]Skill{
			{AllSkills["Aim"]},
			{AllSkills["Larceny"]},
			{AllSkills["Integrity"]},
			{AllSkills["Technology"]},
		},
		Edges: []Edge{
			AllEdges["Always Prepared"],
			AllEdges["Covert"],
			AllEdges["Hair Trigger Reflexes"],
			AllEdges["Small Unit Tactics"],
			AllEdges["Sniper"],
			AllEdges["Wealth"],
		},
	},
	"Adventurer": {
		Name:        "Adventurer",
		Concept:     "",
		Connections: "High-risk Hobbyists (Divers, Mountain Climbers, Stunt Drivers, etc.),	Bomb Disposal Experts, Travel Enthusiasts",
		Skills: [][]Skill{
			{AllSkills["Aim"]},
			{AllSkills["Athletics"]},
			{AllSkills["Pilot"]},
			{AllSkills["Survival"]},
		},
	},
}

var PathsBySplat = map[Splat]map[string][]string{
	Psion: {
		"Origin": {"9"},
	},
}

var AllSkills = map[string]Skill{
	"Aim":          {Name: "Aim"},
	"Athletics":    {Name: "Athletics"},
	"Close Combat": {Name: "Close Combat"},
	"Command":      {Name: "Command"},
	"Culture":      {Name: "Culture"},
	"Empathy":      {Name: "Empathy"},
	"Enigmas":      {Name: "Enigmas"},
	"Humanities":   {Name: "Humanities"},
	"Integrity":    {Name: "Integrity"},
	"Larceny":      {Name: "Larceny"},
	"Medicine":     {Name: "Medicine"},
	"Persuasion":   {Name: "Persuasion"},
	"Pilot":        {Name: "Pilot"},
	"Science":      {Name: "Science"},
	"Survival":     {Name: "Survival"},
	"Technology":   {Name: "Technology"},
}

var AllEdges = map[string]Edge{
	"Always Prepared":       {Name: "Always Prepared"},
	"Covert":                {Name: "Covert"},
	"Hair Trigger Reflexes": {Name: "Hair Trigger Reflexes"},
	"Small Unit Tactics":    {Name: "Small Unit Tactics"},
	"Sniper":                {Name: "Sniper"},
	"Wealth":                {Name: "Wealth"},
}
