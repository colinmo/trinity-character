package main

import "sort"

var AllPaths = map[string]Path{
	"9": {
		Name:        "9",
		Concept:     "",
		Connections: "FBI Agent, Safe House Owner, Weapons Dealer, Lab Worker",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Larcent"}, {"Integrity"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Always Prepared", "Covert", "Hair Trigger Reflexes", "Small Unit Tactics", "Sniper", "Wealth"}),
	},
	"Adventurer": {
		Name:         "Adventurer",
		Concept:      "",
		Connections:  "High-risk Hobbyists (Divers, Mountain Climbers, Stunt Drivers, etc.),	Bomb Disposal Experts, Travel Enthusiasts",
		Skills:       returnSkillsNamed([][]string{{"Aim"}, {"Athletics"}, {"Pilot"}, {"Survival"}}),
		Edges:        returnEdgesNamed([]string{"Always Prepared", "Covert", "Hair Trigger Reflex", "Small Unit Tactics", "Sniper", "Wealth"}),
		GiftKeyAttr:  "Might",
		GiftKeySkill: []string{"Pilot", "Survival"},
	},
	"Life of Privilege": {
		Name:         "Life of Privilege",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Command"}, {"Culture"}, {"Integrity"}, {"Persuasion"}}),
		Edges:        returnEdgesNamed([]string{"Fame", "Patron", "Skilled Liar", "Wealth"}),
		Connections:  "School Alumni, College Club Membership, Local Political Affiliates",
		GiftKeyAttr:  "Presence",
		GiftKeySkill: []string{"Command", "Persuasion"},
	},
	"Military Brat": {
		Name:         "Military Brat",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Command"}, {"Enigmas"}, {"Integrity"}, {"Technology"}}),
		Edges:        returnEdgesNamed([]string{"Adrenaline Spike", "Demolitions Training", "Forceful Martial Arts", "Free Running", "Precise Martial Arts", "Sniper", "Danger Sense", "Fast Draw", "Iron Will", "Patron", "Small Unit Tactics"}),
		Connections:  "Past Teacher, Military Commander, Steadfast Friend",
		GiftKeyAttr:  "Resolve",
		GiftKeySkill: []string{"Integrity", "Technology"},
	},
	"Street Rat": {
		Name:         "Street Rat",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Athletics"}, {"Enigmas"}, {"Larceny"}, {"Survival"}}),
		Edges:        returnEdgesNamed([]string{"Adrenaline Spike", "Alternate Identity", "Always Prepared", "Danger Sense", "Hair Trigger Reflexes", "Hardy", "Ms. Fix-It", "Tough Cookie"}),
		Connections:  "Street Gangs, Street Mentor, Helpful Family Member, Store Clerks",
		GiftKeyAttr:  "Cunning",
		GiftKeySkill: []string{"Larceny", "Survival"},
	},
	"Suburbia": {
		Name:         "Suburbia",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Culture"}, {"Empathy"}, {"Humanities"}, {"Technology"}}),
		Edges:        returnEdgesNamed([]string{"Artistic Talent", "Big Hearted", "Library", "Patron", "Wealth"}),
		Connections:  "Favorite Professor, Neighbor Friend, Influential Teacher",
		GiftKeyAttr:  "Manipulation",
		GiftKeySkill: []string{"Culture", "Empathy"},
	},
	"Survivalist": {
		Name:         "Survivalist",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Aim"}, {"Close Combat"}, {"Medicine"}, {"Survival"}}),
		Edges:        returnEdgesNamed([]string{"Always Prepared", "Animal Ken", "Covert", "Direction Sense", "Hardy", "Iron Will", "Keen Sense", "Swift"}),
		Connections:  "Park Ranger, Conspiracy Groups, RV Neighborhood",
		GiftKeyAttr:  "Stamina",
		GiftKeySkill: []string{"Medicing", "Survival"},
	},
	"Psi: Oceanian": {
		Name:        "Psi: Oceanian",
		Concept:     "",
		Skills:      returnSkillsNamed([][]string{{"Culture"}, {"Empathy"}, {"Survival"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Amphibious or Aquatic Conversion", "Artistic Talent", "Big Hearted", "Hardy", "Increased Tolerance", "Library"}),
		Connections: "Artists or Scientists, Boat or Submarine Pilot, Genetic Engineer, Undersea Miner",
	},
	"Charismatic Leader": {
		Name:         "Charismatic Leader",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Command"}, {"Empathy"}, {"Humanities"}, {"Persuasion"}}),
		Edges:        returnEdgesNamed([]string{"Fame", "Iron Will", "Skilled Liar", "Striking", "Wealth"}),
		Connections:  "Corporate Board, Megachurch, Political Allies",
		GiftKeyAttr:  "Manipulation",
		GiftKeySkill: []string{"Command", "Humanities"},
	},
	"Combat Specialist": {
		Name:         "Combat Specialist",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Aim"}, {"Athletics"}, {"Close Combat"}, {"Integrity"}}),
		Edges:        returnEdgesNamed([]string{"Alternate Identity", "Demolitions Training", "Forceful Martial Arts", "Free Running", "Precise Martial Arts", "Sniper", "Armor Expert", "Breath Control", "Fast Draw", "Hair Trigger Reflexes", "Small Unit Tactics", "Trick Shooter", "Weak Spots"}),
		Connections:  "Military Unit, Police Officers, Training Master",
		GiftKeyAttr:  "Might",
		GiftKeySkill: []string{"Aim", "Close Combat"},
	},
	"Detective": {
		Name:         "Detective",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Aim"}, {"Enigmas"}, {"Integrity"}, {"Persuasion"}}),
		Edges:        returnEdgesNamed([]string{"Alternate Identity", "Demolitions Training", "Forceful Martial Arts", "Free Running", "Precise Martial Arts", "Sniper", "Fast Draw", "Library", "Photographic Memory", "Swift", "Tough Cookie"}),
		Connections:  "Police Officers, Paid Informant, New Reporter, Friendly Neighborhood Watch",
		GiftKeyAttr:  "Cunning",
		GiftKeySkill: []string{"Aim", "Enigmas"},
	},
	"Medical Practitioner": {
		Name:         "Medical Practitioner",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Empathy"}, {"Medicine"}, {"Science"}, {"Survival"}}),
		Edges:        returnEdgesNamed([]string{"Always Prepared", "Ambidextrous", "Big Hearted", "Iron Will", "Keen Sense", "Library", "Wealth"}),
		Connections:  "Surgeon, Pharmacists, Thankful Patient, EMTs",
		GiftKeyAttr:  "Resolve",
		GiftKeySkill: []string{"Medicine", "Science"},
	},
	"Pilot": {
		Name:         "Pilot",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Aim"}, {"Close Combat"}, {"Pilot"}, {"Technology"}}),
		Edges:        returnEdgesNamed([]string{"Ambidextrous", "Cool Under Fire", "Demolitions Training", "Direction Sense", "Hair Trigger Reflexes", "Ms. Fix-It", "Patron", "Tough Cookie"}),
		Connections:  "Important Client, Criminal Organization, Indebted Passenger",
		GiftKeyAttr:  "Dexterity",
		GiftKeySkill: []string{"Pilot", "Technology"},
	},
	"The Sneak": {
		Name:         "The Sneak",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Athletics"}, {"Enigmas"}, {"Larceny"}, {"Technology"}}),
		Edges:        returnEdgesNamed([]string{"Adrenaline Spike", "Alternate Identity", "Covert", "Free Running", "Photographic Memory", "Skilled Liar"}),
		Connections:  "Criminal Organization, Best Friend, Police Insider",
		GiftKeyAttr:  "Composure",
		GiftKeySkill: []string{"Athletics", "Larceny"},
	},
	"Technology Expert": {
		Name:         "Technology Expert",
		Concept:      "",
		Skills:       returnSkillsNamed([][]string{{"Culture"}, {"Enigmas"}, {"Science"}, {"Technology"}}),
		Edges:        returnEdgesNamed([]string{"Demolitions Training", "Library", "Lightning Calculator", "Ms. Fixit", "Patron", "Weak Spots", "Swift"}),
		Connections:  "Chop Shop Worker, Research Scientists, Machinist Friend",
		GiftKeyAttr:  "Intellect",
		GiftKeySkill: []string{"Science", "Technology"},
	},
	"Psi: Off-Earth Colonist": {
		Name:        "Psi: Off-Earth Colonist",
		Concept:     "",
		Skills:      returnSkillsNamed([][]string{{"Athletics"}, {"Integrity"}, {"Survival"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Animal Ken", "Danger Sense", "Direction Sense", "Hardy", "Ms. Fix-It", "Pioneer Spirit", "Superior Trait", "Weak Spots"}),
		Connections: "Corporate Sponsor, Interstellar Explorer, Survivalist, Xenobiologist",
	},
	"Psi: Spacer": {
		Name:        "Psi: Spacer",
		Concept:     "",
		Skills:      returnSkillsNamed([][]string{{"Enigmas"}, {"Pilot"}, {"Science"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Breath Control", "Endurance", "Hardy", "Low Gravity Adaptation", "Ms. Fix-It", "Variable Gravity Training", "Weak Spots"}),
		Connections: "Asteroid Miner, Interstellar Explorer, Off-Earth Colonist, Upeo Wa Macho Member",
	},
	"Psi: Space Military": {
		Name:        "Psi: Space Military",
		Concept:     "",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Command"}, {"Pilot"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Adrenaline Spike", "Breath Control", "Danger Sense", "Hair-Trigger Reflexes", "Hardy", "Low Gravity Adaptation", "Small Unit Tactics", "Variable Gravity Training", "Variable Gravity Combat Training"}),
		Connections: "Armorer, Off-Earth Colony, Spacer Captain, UNMC",
	},
	"Æon Socity": {
		Name:        "Æon Socity",
		Concept:     "",
		Connections: "High political figure, Military Advisor, Large Charity Fund Manager",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Close Combat"}, {"Enigmas"}, {"Pilot"}}),
		Edges:       returnEdgesNamed([]string{"Always Prepared", "Direction Sense", "Artifact", "Library", "Wealth"}),
	},
	"Archangel": {
		Name:        "Archangel",
		Concept:     "",
		Connections: "Pro Bono Lawyer, Witness Protection Officer, Homeland Security Officer, Criminal with a Heart of Gold, Hactivist",
		Skills:      returnSkillsNamed([][]string{{"Close Combat"}, {"Empathy"}, {"Integrity"}, {"Persuasion"}}),
		Edges:       returnEdgesNamed([]string{"Adrenaline Spike", "Big Hearted", "Endurance", "Iron Will", "Patron", "Skilled Liar", "Speed Reading"}),
	},
	"The Global Cartography Initiative": {
		Name:        "The Global Cartography Initiative",
		Concept:     "",
		Connections: "Black Market Artifact Dealer, Smuggler, Museum Curator, Border Guard, Journalist, Mercenary, Pirate, Pirate Hunter",
		Skills:      returnSkillsNamed([][]string{{"Enigmas"}, {"Humanities"}, {"Larceny"}, {"Survival"}}),
		Edges:       returnEdgesNamed([]string{"Artifact", "Direction Sense", "Library", "Patron"}),
	},
	"Neptune Foundation": {
		Name:        "Neptune Foundation",
		Concept:     "",
		Connections: "Aid Worker, Emergency Services, ER Doctor, Free Clinic Volunteer, Local Government Representative",
		Skills:      returnSkillsNamed([][]string{{"Command"}, {"Integrity"}, {"Medicine"}, {"Persuasion"}}),
		Edges:       returnEdgesNamed([]string{"Fame", "Iron Will", "Keen Sense", "Patron", "Photographic Memory", "Superior Trait"}),
	},
	"Pharaoh's Lightkeepers": {
		Name:        "Pharaoh's Lightkeepers",
		Concept:     "",
		Connections: "Journalists, Military Personnel, Other Lightkeeper Teams, Police Officers, and all manner of ordinary citizens",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Close Combat"}, {"Enigmas"}, {"Pilot"}}),
		Edges:       returnEdgesNamed([]string{"Artifact", "Danger Sense", "Library", "Skilled Liar", "Small Unit Tactics", "Sniper"}),
	},
	"Alert Status 1": {
		Name:        "Alert Status 1",
		Concept:     "",
		Connections: "Committee Member, National Intelligence Director, Friendly Agent of a Rival Nation",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Enigmas"}, {"Persuasion"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Alternate Identity", "Armor Expert", "Cool Under Fire", "Covert", "Direction Sense", "Sniper", "Trick Shooter"}),
	},
	"La Révolte Éclatante": {
		Name:        "La Révolte Éclatante",
		Concept:     "",
		Connections: "Idealistic Priests, Labor Organizers, Medical Relief Personnel, Street Gangs, Violent Anarchists",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Medicine"}, {"Pilot"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Alternate Identity", "Cool Under Fire", "Demolitions Training", "Safe House", "Small Unit Tactics", "Swift", "Tough Cookie", "Weak Spots"}),
	},
	"Les Fantômes": {
		Name:        "Les Fantômes",
		Concept:     "",
		Connections: "Fence, Forger, Grateful Museum Official, Grudgingly Respectful Interpol Agent",
		Skills:      returnSkillsNamed([][]string{{"Athletics"}, {"Culture"}, {"Larceny"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Covert", "Free Running", "Safe House", "Skilled Liar", "Wealth"}),
	},
	"National Office of Emergency Research": {
		Name:        "National Office of Emergency Research",
		Concept:     "",
		Connections: "Anonymous Online Source, Off-Record Inside Informant, Paraphysical Research Study Group, UFO Witness",
		Skills:      returnSkillsNamed([][]string{{"Command"}, {"Enigmas"}, {"Humanities"}, {"Persuasion"}}),
		Edges:       returnEdgesNamed([]string{"Always Prepared", "Artifact", "Covert", "Patron", "Small Unit Tactics", "Speed Reading"}),
	},
	"The Theseus Club": {
		Name:        "The Theseus Club",
		Concept:     "",
		Connections: "FBI Agent, Local Hunting Club President, Wealthy Do-Gooder",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Athletics"}, {"Larceny"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Alternate Identity", "Always Prepared", "Endurance", "Danger Sense", "Demolitions Expert", "Small Unit Tactics", "Trick Shooter"}),
	},
	"Transcendent Alliance": {
		Name:        "Transcendent Alliance",
		Concept:     "",
		Connections: "Cutting Edge Scientists, Grey Market Pharmaceutical Manufacturers, International Smugglers, Skilled Programmers",
		Skills:      returnSkillsNamed([][]string{{"Culture"}, {"Medicine"}, {"Science"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Lightning Calculator", "Ms. Fix-It", "Photographic Memory", "Superior Trait", "Weak Spots", "Wealth"}),
	},
	"Triton Foundation": {
		Name:        "Triton Foundation",
		Concept:     "",
		Connections: "Medical Researcher, Famous Surgeon, President of a Charity, Local Public Leader, Dean of a Research College",
		Skills:      returnSkillsNamed([][]string{{"Enigmas"}, {"Medicine"}, {"Persuasion"}, {"Science"}}),
		Edges:       returnEdgesNamed([]string{"Ambidextrous", "Big Hearted", "Iron Will", "Library", "Superior Trait", "Wealth"}),
	},
	"Psi: Freelance Psion": {
		Name:        "Psi: Freelance Psion",
		Concept:     "",
		Connections: "Æon Trinity, Criminal Organization, Psi Order, Thankful Client",
		Skills:      returnSkillsNamed([][]string{{"Empathy"}, {"Integrity"}, {"Larceny"}, {"Persuasion"}}),
		Edges:       returnEdgesNamed([]string{"Adrenaline Spike", "Always Prepared", "Covert", "Danger Sense", "Enhanced Attunement", "Favored Mode", "Inner Reserve", "Patron"}),
	},
	"Psi: Nihonjin Agent": {
		Name:        "Psi: Nihonjin Agent",
		Concept:     "",
		Connections: "Criminal Organization, Double Agent, Government Official, Nippon Intelligence Agency",
		Skills:      returnSkillsNamed([][]string{{"Athletics"}, {"Enigmas"}, {"Larceny"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Alternative Identity", "Computer Enhancement", "Covert", "Cyberware Access", "Danger Sense", "Skilled Liar", "Superior Trait"}),
	},
	"Psi: Æsculapian Order": {
		Name:        "Psi: Æsculapian Order",
		Concept:     "",
		Connections: "Æon Trinity, Emergency Service Personnel, Health Ministry, Medical Specialist",
		Skills:      returnSkillsNamed([][]string{{"Athletics"}, {"Empathy"}, {"Medicine"}, {"Science"}}),
		Edges:       returnEdgesNamed([]string{"Adrenaline Spike", "Danger Sense", "Enhanced Attunement", "Favored Mode", "Hardy", "Keen Sense", "Library", "Rapid Response Protocols", "Trained Memory"}),
	},
	"Psi: ISRA": {
		Name:        "Psi: ISRA",
		Concept:     "",
		Connections: "Æon Trinity, Local Community Organizer, Respected Priest, Seemingly Random Individual Clairsentience Indicated To Be Significant",
		Skills:      returnSkillsNamed([][]string{{"Culture"}, {"Empathy"}, {"Humanities"}, {"Integrity"}}),
		Edges:       returnEdgesNamed([]string{"Artistic Talent", "Danger Sense", "Direction Sense", "Enhanced Attunement", "Favored Mode", "Keen Sense", "Mysterious Aid", "Patron", "Weak Spots"}),
	},
	"Psi: Legion": {
		Name:        "Psi: Legion",
		Concept:     "",
		Connections: "Local Militia, Mercenary Company, Military Contractor, UN Official",
		Skills:      returnSkillsNamed([][]string{{"Aim"}, {"Athletics"}, {"Close Combat"}, {"Command"}}),
		Edges:       returnEdgesNamed([]string{"Always Prepared", "Enhanced Attunement", "Favored Mode", "Hardy", "Iron Will", "Legion Armory", "Small Unit Tactics"}),
	},
	"Psi: Ministry": {
		Name:        "Psi: Ministry",
		Concept:     "",
		Connections: "Government Official, Local Religious Leader, Noetics Expert, Politician",
		Skills:      returnSkillsNamed([][]string{{"Culture"}, {"Empathy"}, {"Enigmas"}, {"Integrity"}}),
		Edges:       returnEdgesNamed([]string{"Covert", "Dual Minded", "Favored Mode", "Iron Will", "Subtle Influence", "Telepathic Resistance"}),
		// "TEdges": ["Covert", "Danger Sense", "Enhanced Attunement", "Favored Mode", "Iron Will", "Skilled Liar", "Subtle Influence", "Trained Memory"]
	},
	"Psi: Norça": {
		Name:        "Psi: Norça",
		Concept:     "",
		Connections: "Ecoterrorist, Ecoscientist, Humanitarian Agency, Intelligence Agency",
		Skills:      returnSkillsNamed([][]string{{"Athletics"}, {"Empathy"}, {"Science"}, {"Survival"}}),
		Edges:       returnEdgesNamed([]string{"Ambidextrous", "Enhanced Attunement", "Favored Mode", "Alternate Identity", "Hardy", "Impersonation Training", "Keen Sense", "Safe House", "Superior Trait"}),
	},
	"Psi: Orgotek": {
		Name:        "Psi: Orgotek",
		Concept:     "",
		Connections: "Biotechnician, FSA Regulator, High-Flying CEO, Start-Up Business",
		Skills:      returnSkillsNamed([][]string{{"Culture"}, {"Persuasion"}, {"Science"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Artifact", "Bioware Access", "Enhanced Attunement", "Favored Mode", "Increased Tolerance", "Wealth"}),
	},
	"Psi: Upeo wa Macho": {
		Name:        "Psi: Upeo wa Macho",
		Concept:     "",
		Connections: "Bonded Witness, Interstellar Explorer, Leviathan Captain, Off-Earth Colony",
		Skills:      returnSkillsNamed([][]string{{"Integrity"}, {"Pilot"}, {"Survival"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Direction Sense", "Emergency Transit", "Enhanced Attunement", "Favored Mode", "Hardy", "Iron Will", "Photographic Memory"}),
	},
	"Psi: Æon Trinity": {
		Name:        "Psi: Æon Trinity",
		Concept:     "",
		Connections: "Æon Council, Humanitarian Group, Political Allies, Psi Order",
		Skills:      returnSkillsNamed([][]string{{"Humanities"}, {"Persuasion"}, {"Science"}, {"Technology"}}),
		Edges:       returnEdgesNamed([]string{"Fame", "Far-Reaching Influence", "Library", "Superior Trait", "Wealth"}),
	},
	"Psi: Æon Trinity Section Minerva": {
		Name:        "Psi: Æon Trinity Section Minerva",
		Concept:     "",
		Connections: "Æon Council, Criminal Organization, Mercenary Company, Spy Agency",
		Skills:      returnSkillsNamed([][]string{{"Empathy"}, {"Integrity"}, {"Larceny"}, {"Persuasion"}}),
		Edges:       returnEdgesNamed([]string{"Alternate Identity", "Always Prepared", "Covert", "Danger Sense", "Iron Will", "Self-Reliance"}),
	},
}

var PathsBySplat = map[Splat]map[string][]string{
	Talent: {
		"Origin":  {"Adventurer", "Life of Privilege", "Military Brat", "Street Rat", "Suburbia", "Survivalist"},
		"Role":    {"Charismatic Leader", "Combat Specialist", "Detective", "Medical Practitioner", "Pilot", "The Sneak", "Technology Expert"},
		"Society": {"9", "Æon Socity", "Archangel", "The Global Cartography Initiative", "The Neptune Foundation", "Pharaoh's Lightkeepers", "Alert Status 1", "La Révolte Éclatante", "Les Fantômes", "National Office of Emergency Research", "The Theseus Club", "The Transcendent Alliance", "Triton Foundation"},
	},
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

func returnSkillsNamed(names [][]string) [][]Skill {
	toReturn := [][]Skill{}
	for _, nameGroup := range names {
		mep := []Skill{}
		for _, name := range nameGroup {
			mep = append(mep, AllSkills[name])
		}
		toReturn = append(toReturn, mep)
	}
	return toReturn
}
func returnEdgesNamed(names []string) []Edge {
	toReturn := []Edge{}
	for _, name := range names {
		toReturn = append(toReturn, AllEdges[name])
	}
	return toReturn
}

func returnAlphaSkill() []string {
	mk := make([]string, len(AllSkills))
	i := 0
	for k := range AllSkills {
		mk[i] = k
		i++
	}
	sort.Strings(mk)
	return mk
}
