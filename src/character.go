package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
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

type Dot2Button struct {
	Max     int
	Min     int
	Value   int
	Buttons []*tappableIcon
}

func (d *Dot2Button) SetValue(value int) {
	d.Value = value
	// Set button appearances
}

func (d *Dot2Button) ToCanvas() *fyne.Container {
	x := fyne.NewContainerWithLayout(
		layout.NewHBoxLayout(),
	)
	for _, y := range d.Buttons {
		x.Objects = append(x.Objects, y)
	}
	return x
}

func MakeDotButtons(max int, min int) Dot2Button {
	var d Dot2Button
	d.Max = max
	d.Min = min
	d.Value = min

	d.Buttons = make([]*tappableIcon, max)
	for x := range d.Buttons {
		z := x
		icon := theme.RadioButtonIcon()
		if x < d.Value {
			icon = theme.RadioButtonCheckedIcon()
		}
		d.Buttons[x] = newTappableIcon(icon, func() {
			thisValue := z + 1
			// @todo: Check if this is a valid value
			if thisValue < d.Min || thisValue == d.Value {
				fmt.Printf("Can't go below min")
				d.Value = d.Min
			} else {
				d.Value = thisValue
			}
			for mep := 0; mep < d.Value; mep++ {
				d.Buttons[mep].Icon.Resource = theme.RadioButtonCheckedIcon()
				d.Buttons[mep].Refresh()
			}
			for mep := d.Value; mep < d.Max; mep++ {
				d.Buttons[mep].Icon.Resource = theme.RadioButtonIcon()
				d.Buttons[mep].Refresh()
			}
			fmt.Printf("Tapped %d\n", z+1)
		})
	}
	return d
}

func MakeEditableCharacterSheet() *fyne.Container {
	mep0 := MakeDotButtons(5, 1)
	mep1 := MakeDotButtons(5, 1)
	mep2 := MakeDotButtons(5, 1)
	attributesLine := container.New(
		layout.NewHBoxLayout(),
		container.New(
			layout.NewVBoxLayout(),
			canvas.NewText("Force", color.Black),
			canvas.NewText("Finesse", color.Black),
			canvas.NewText("Resilience", color.Black),
		),
		container.New(
			layout.NewFormLayout(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Intellect", color.Black)),
			mep0.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Cunning", color.Black)),
			mep1.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Resolve", color.Black)),
			mep2.ToCanvas(),
		),
		container.New(
			layout.NewFormLayout(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Might", color.Black)),
			mep0.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Dexterity", color.Black)),
			mep1.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Stamina", color.Black)),
			mep2.ToCanvas(),
		),
		container.New(
			layout.NewFormLayout(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Presence", color.Black)),
			mep0.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Manipulation", color.Black)),
			mep1.ToCanvas(),
			container.New(layout.NewHBoxLayout(), canvas.NewText("Composure", color.Black)),
			mep2.ToCanvas(),
		),
	)
	skillsLine := container.New(
		layout.NewGridLayoutWithColumns(2),
		container.New(
			layout.NewFormLayout(),
			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(layout.NewHBoxLayout(), canvas.NewText("Aim", color.Black))),
			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(
					layout.NewHBoxLayout(),
					layout.NewSpacer(),
					container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
				)),

			container.New(layout.NewHBoxLayout(), canvas.NewText("Athletics", color.Black)),
			container.New(
				layout.NewHBoxLayout(),
				layout.NewSpacer(),
				container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
			),

			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(layout.NewHBoxLayout(), canvas.NewText("Close Combat", color.Black))),
			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(
					layout.NewHBoxLayout(),
					layout.NewSpacer(),
					container.New(layout.NewCenterLayout(), mep0.ToCanvas())),
			),
		),
		container.New(
			layout.NewFormLayout(),
			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(layout.NewHBoxLayout(), canvas.NewText("Integrity", color.Black))),
			container.New(layout.NewMaxLayout(), canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(layout.NewHBoxLayout(), layout.NewSpacer(), container.New(layout.NewCenterLayout(), mep0.ToCanvas()))),
		),
	)
	return container.New(
		layout.NewVBoxLayout(),
		// Basic Info
		container.New(
			layout.NewGridLayoutWithColumns(2),
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Name"),
				widget.NewEntry(),
				widget.NewLabel("Player"),
				widget.NewEntry(),
				widget.NewLabel("Concept"),
				widget.NewEntry(),
			),
			container.New(
				layout.NewGridLayoutWithColumns(3),
				widget.NewLabel("Origin Path"),
				widget.NewSelect([]string{"Path1", "Path2"}, func(option string) { fmt.Printf("String %s\n", option) }),
				container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
				widget.NewLabel("Role Path"),
				widget.NewSelect([]string{"Path1", "Path2"}, func(option string) { fmt.Printf("String %s\n", option) }),
				container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
				widget.NewLabel("Allegiance Path"),
				widget.NewSelect([]string{"Path1", "Path2"}, func(option string) { fmt.Printf("String %s\n", option) }),
				container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
			),
		),
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Moment of Inspiration"),
			widget.NewEntry(),
		),
		// Aspirations
		// Skills
		widget.NewLabel("SKILLS"),
		skillsLine,
		// Skill Tricks
		// Attributes
		widget.NewLabel("ATTRIBUTES"),
		attributesLine,
		// Edges
		// Health
		// Chartype Specific
	)
}
