package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Approach int64

const (
	Force Approach = iota
	Finesse
	Resilience
)

type Skill struct {
	Name         string
	Blurb        string
	Specialities []string
	SkillTricks  []string
}

type PreRequisite struct {
	LinkedThing string
	MinDots     int
	MaxDots     int
}
type Edge struct {
	Name          string
	Blurb         string
	PreRequisites []PreRequisite
	Group         string
	Levels        []string
	Multiple      bool
}
type Gift struct {
	Name  string
	Blurb string
}

type Path struct {
	Name        string
	Concept     string
	Connections string
	Skills      []Skill
	Edges       []Edge
}

type Moment struct {
	Name                string
	AssociatedAttribute string
}

type TrinityCharacter struct {
	Name        string
	Concept     string
	Type        string
	Aspirations struct {
		ShortTerm []string
		LongTerm  []string
	}
	MomentsOfInspiration []Moment
	Paths                struct {
		Origin  Path
		Role    Path
		Society Path
		Others  []Path
	}
	Skills []struct {
		Skill Skill
		Dots  int
	}
	Edges []struct {
		Edge Edge
		Dots int
	}
	Attributes struct {
		Might        int
		Dexterity    int
		Stamina      int
		Intellect    int
		Cunning      int
		Resolve      int
		Presence     int
		Manipulation int
		Composure    int
	}
	FavouredApproach Approach
	Facets           struct {
		Intuitive   int
		Reflective  int
		Destructive int
	}
	Health struct {
		Total      int
		Bashing    int
		Lethal     int
		Aggravated int
	}
	Injuries []struct {
		Label   string
		Penalty int
		Marked  bool
	}
	Defense     int
	Armor       int
	Experiences int
	Gifts       []Gift
}

type Aptitude struct {
	Name  string
	Basic string
	Modes struct {
		Label string
		Dots  int
	}
	Proxy string
}

type PSI struct {
	Dots            int
	PrimaryAptitude Aptitude
	AuxillaryModes  struct {
		Label string
		Dots  int
	}
}

func characterWindow(app *fyne.App, charName string) fyne.Window {
	mainWindow := (*app).NewWindow(charName)
	mainWindow.SetContent(widget.NewLabel("Hello"))
	return mainWindow

}
