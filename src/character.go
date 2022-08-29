package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

type Splat int64

const (
	Talent Splat = iota
	Psion
	Aberrant
	Adventure
	Aether
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
	Name         string
	Concept      string
	Connections  string
	Skills       [][]Skill
	Edges        []Edge
	GiftKeySkill []string
	GiftKeyAttr  string
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
	x := container.New(
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
			layout.NewVBoxLayout(),
			container.New(
				layout.NewMaxLayout(),
				canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(
					layout.NewHBoxLayout(),
					canvas.NewText("Aim", color.Black),
					layout.NewSpacer(),
					container.New(layout.NewCenterLayout(), mep0.ToCanvas())),
			),
			container.New(
				layout.NewHBoxLayout(),
				canvas.NewText("Athletics", color.Black),
				layout.NewSpacer(),
				container.New(layout.NewCenterLayout(), mep0.ToCanvas()),
			),
			container.New(
				layout.NewMaxLayout(),
				canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(
					layout.NewHBoxLayout(),
					canvas.NewText("Close Combat", color.Black),
					layout.NewSpacer(),
					container.New(layout.NewCenterLayout(), mep0.ToCanvas())),
			),
		),
		container.New(
			layout.NewVBoxLayout(),
			container.New(
				layout.NewMaxLayout(),
				canvas.NewRectangle(color.NRGBA{R: 200, G: 100, B: 0, A: 33}),
				container.New(
					layout.NewHBoxLayout(),
					canvas.NewText("Integrity", color.Black),
					layout.NewSpacer(),
					container.New(layout.NewCenterLayout(), mep0.ToCanvas())),
			),
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

func MakeCharacterCreationScreen(splat Splat, window *fyne.Window) *fyne.Container {
	// Character variables
	// Set first so they can be adjusted during the creation
	charConcept := map[string]*widget.Entry{}
	charConcept["Name"] = widget.NewEntry()
	charConcept["Player"] = widget.NewEntry()
	charConcept["Concept"] = widget.NewEntry()
	charConcept["Short Aspiration 1"] = widget.NewEntry()
	charConcept["Short Aspiration 2"] = widget.NewEntry()
	charConcept["Long Aspiration"] = widget.NewMultiLineEntry()
	skillEntries := map[string]*dotSelector{}
	skillLabels := map[string]*widget.Label{}
	attributes := map[string]*dotSelector{
		"Intellect":    newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Cunning":      newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Resolve":      newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Might":        newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Dexterity":    newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Stamina":      newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Presence":     newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Manipulation": newDotSelector(1, 5, 1), // No more than 5 at char creation
		"Composure":    newDotSelector(1, 5, 1), // No more than 5 at char creation
	}
	favApproach := map[string]*widget.Check{
		"Force":      widget.NewCheck("Force", func(meOn bool) {}),
		"Finesse":    widget.NewCheck("Finesse", func(meOn bool) {}),
		"Resilience": widget.NewCheck("Resilience", func(meOn bool) {}),
	}
	for i := range favApproach {
		var bob = i
		favApproach[i].OnChanged = func(meOn bool) {
			if meOn {
				for i, x := range favApproach {
					if x.Text != bob {
						favApproach[i].SetChecked(false)
					}
				}
			}
		}
	}
	// Interface variables
	// Set first so they can be adjusted during the creation
	conceptTab := container.NewTabItemWithIcon(
		"Concept",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	pathsTab := container.NewTabItemWithIcon(
		"Paths",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	skillsTab := container.NewTabItemWithIcon(
		"Skills",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	attributesTab := container.NewTabItemWithIcon(
		"Attributes",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	templateTab := container.NewTabItemWithIcon(
		"Template",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	finishingTab := container.NewTabItemWithIcon(
		"Finishing touches",
		theme.CheckButtonIcon(),
		container.NewWithoutLayout(),
	)
	header := "Trinity"
	// Step 1: Concept
	baseStep1 := []fyne.CanvasObject{
		widget.NewLabel("Name"),
		charConcept["Name"],
		widget.NewLabel("Player"),
		charConcept["Player"],
		widget.NewLabel("Concept"),
		charConcept["Concept"],
		widget.NewLabel("Aspiration (Short)"),
		charConcept["Short Aspiration 1"],
		widget.NewLabel("Aspiration (Short)"),
		charConcept["Short Aspiration 2"],
		widget.NewLabel("Aspiration (Long)"),
		charConcept["Long Aspiration"],
	}

	// Step 2: Paths
	// Get the splat specific configurations like Paths
	baseStep2 := append(pathSelectorFor("Origin", splat, window), pathSelectorFor("Role", splat, window)...)
	baseStep2 = append(baseStep2, pathSelectorFor("Society", splat, window)...)

	// Step 3: Skills
	baseStep3a := container.New(layout.NewFormLayout())
	baseStep3b := container.New(layout.NewFormLayout())
	half := len(AllSkills) / 2
	count := 0
	for _, key := range returnAlphaSkill() {
		skillEntries[key] = newDotSelector(0, 5, 0)
		skillLabels[key] = widget.NewLabel(AllSkills[key].Name)
		if count < half {
			baseStep3a.Objects = append(
				baseStep3a.Objects,
				skillLabels[key],
				skillEntries[key],
			)
		} else {
			baseStep3b.Objects = append(
				baseStep3b.Objects,
				skillLabels[key],
				skillEntries[key],
			)
		}
		count++
	}

	skillLabels["Aim"].Text = "Aim (+2)"

	switch splat {
	case Psion:
		header = "Psion"
		baseStep1 = append(baseStep1,
			widget.NewLabel("Psi Order"),
			widget.NewEntry())
	case Aberrant:
		header = "Aberrant"
	case Aether:
		header = "Aether"
	case Adventure:
		header = "Adventure!"
	}

	// Step 4: Attributes
	attributesHelp := widget.NewLabel("Characters begin with a single Attribute dot in each of their nine Attributes. Players distribute six dots among the three Attributes in their top-ranked Arena, four dots in their middle-ranked, and two dots in the bottom-ranked.")
	attributesHelp.Wrapping = fyne.TextWrapWord
	baseStep4 := container.NewGridWithColumns(4,
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel(""),
			favApproach["Force"],
			widget.NewLabel(""),
			favApproach["Finesse"],
			widget.NewLabel(""),
			favApproach["Resilience"],
		),
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Intellect"),
			attributes["Intellect"],
			widget.NewLabel("Cunning"),
			attributes["Cunning"],
			widget.NewLabel("Resolve"),
			attributes["Resolve"],
			widget.NewLabel("                              "),
			widget.NewLabel("     "),
		),
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Might"),
			attributes["Might"],
			widget.NewLabel("Dexterity"),
			attributes["Dexterity"],
			widget.NewLabel("Stamina"),
			attributes["Stamina"],
			widget.NewLabel("                              "),
			widget.NewLabel("     "),
		),
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Presence"),
			attributes["Presence"],
			widget.NewLabel("Manipulation"),
			attributes["Manipulation"],
			widget.NewLabel("Composure"),
			attributes["Composure"],
			widget.NewLabel("                              "),
			widget.NewLabel("     "),
		),
	)
	// Step 5: Template. This is custom per Splat
	baseStep5 := container.NewMax(widget.NewLabel("Custom template"))
	switch splat {
	case Talent:
		momentOfInspirationText := widget.NewEntry()
		momentOfInspirationAttr := widget.NewSelect([]string{
			"Intellect", "Cunning", "Resolve", "Might", "Dexterity", "Stamina", "Presence", "Manipulation", "Composure",
		}, func(selected string) {})
		gifts := []*widget.Entry{
			widget.NewEntry(),
			widget.NewEntry(),
			widget.NewEntry(),
			widget.NewEntry(),
		}
		giftLabels := []*widget.Label{
			widget.NewLabel("Gift, Origin Path"),
			widget.NewLabel("Gift, Role Path"),
			widget.NewLabel("Gift, Allegiance Path"),
			widget.NewLabel("Gift, Open"),
		}
		facets := map[string]*dotSelector{
			"Intuitive":   newDotSelector(0, 5, 0),
			"Reflective":  newDotSelector(0, 5, 0),
			"Destructive": newDotSelector(0, 5, 0),
		}
		inspiration := widget.NewLabel("1")
		for i := range facets {
			bob := facets[i].OnChanged
			facets[i].OnChanged = func(s string) {
				bob(s)
				fmt.Printf("Facets %d|%d|%d\n", facets["Intuitive"].Val, facets["Reflective"].Val, facets["Destructive"].Val)
				inspiration.SetText(
					fmt.Sprintf(
						"%d", 1+
							((facets["Intuitive"].Val+1)/2)+
							((facets["Reflective"].Val+1)/2)+
							((facets["Destructive"].Val+1)/2)),
				)
			}
		}
		baseStep5 =
			container.New(
				layout.NewFormLayout(),
				widget.NewLabel("Moment of Inspiration"),
				momentOfInspirationText,
				widget.NewLabel("Associated attribute"),
				momentOfInspirationAttr,
				giftLabels[0],
				gifts[0],
				giftLabels[1],
				gifts[1],
				giftLabels[2],
				gifts[2],
				giftLabels[3],
				gifts[3],
				widget.NewLabel("Intuitive"),
				facets["Intuitive"],
				widget.NewLabel("Reflective"),
				facets["Reflective"],
				widget.NewLabel("Destructive"),
				facets["Destructive"],
				widget.NewLabel("Inspiration"),
				inspiration,
			)
	}
	// Step 6: Final touches
	// Fill tabs with content
	conceptTab.Content = container.New(layout.NewFormLayout(), baseStep1...)
	pathsTab.Content = container.New(layout.NewFormLayout(), baseStep2...)
	skillsTab.Content = container.NewGridWithColumns(2, baseStep3a, baseStep3b)
	attributesTab.Content =
		container.NewVBox(
			attributesHelp,
			baseStep4,
		)
	templateTab.Content = container.NewVBox(baseStep5)
	finishingTab.Content = container.NewMax(widget.NewLabel("Bonus trait, 4 points of edges, health, defense"))
	// Present the tabbed interface
	appTabs := container.NewAppTabs(
		conceptTab,
		pathsTab,
		skillsTab,
		attributesTab,
		templateTab,
		finishingTab,
	)
	// Use this to validate the content on leaving
	appTabs.OnUnselected = func(tab *container.TabItem) {
		switch tab.Text {
		case "Concept":
			for _, y := range charConcept {
				if len(y.Text) == 0 {
					conceptTab.Icon = theme.CheckButtonIcon()
					return
				}
			}
			conceptTab.Icon = theme.CheckButtonCheckedIcon()
		}
		fmt.Printf("Unselected %v\n", tab)
	}
	return container.NewVBox(
		container.NewMax(
			widget.NewLabel(header),
		),
		appTabs,
	)
}

func pathSelectorFor(path string, splat Splat, window *fyne.Window) []fyne.CanvasObject {
	availablePaths := map[string]Path{}
	availablePathIndexes := []string{}
	for _, x := range PathsBySplat[splat][path] {
		availablePaths[x] = AllPaths[x]
		availablePathIndexes = append(availablePathIndexes, x)
	}
	skillPrompts := []fyne.CanvasObject{
		container.NewHBox(widget.NewLabel("Skill 1"), newDotSelector(0, 5, 0)),
		container.NewHBox(widget.NewLabel("Skill 2"), newDotSelector(0, 5, 0)),
		container.NewHBox(widget.NewLabel("Skill 3"), newDotSelector(0, 5, 0)),
		container.NewHBox(widget.NewLabel("Skill 4"), newDotSelector(0, 5, 0)),
	}
	edgePrompts := []fyne.CanvasObject{
		container.NewHBox(widget.NewLabel("Edge 1"), newDotSelector(0, 5, 0)),
		container.NewHBox(widget.NewLabel("Edge 2"), newDotSelector(0, 5, 0)),
	}
	pathDescription := widget.NewMultiLineEntry()
	pathDescription.Disable()
	pathSelectBox := widget.NewSelect(
		availablePathIndexes,
		func(changed string) {
			if len(changed) > 0 {
				pathDescription.Text = fmt.Sprintf("%v", availablePaths[changed])
				pathDescription.Refresh()
			}
		})
	selector := []fyne.CanvasObject{
		widget.NewButton(path, func() {
			// Prompt for Path
			dialog.ShowCustomConfirm(
				"Path selector - "+path,
				"Select",
				"Cancel",
				container.NewVBox(
					pathSelectBox,
					pathDescription,
				),
				func(isok bool) {
					if isok {
						thisPath := availablePaths[pathSelectBox.Selected]
						for i, x := range thisPath.Skills {
							if len(x) == 1 {
								skillPrompts[i] = container.NewHBox(widget.NewLabel(x[0].Name), newDotSelector(0, 5, 0))
							} else {
								skillOptions := []string{}
								for _, z := range x {
									skillOptions = append(skillOptions, z.Name)
								}
								skillPrompts[i] = container.NewHBox(widget.NewSelect(skillOptions, func(bob string) {
								}), newDotSelector(0, 5, 0))
							}
							skillPrompts[i].Refresh()
						}
						edgeOptions := []string{}
						for _, x := range thisPath.Edges {
							edgeOptions = append(edgeOptions, x.Name)
						}
						for i := range edgePrompts {
							edgePrompts[i] = container.NewHBox(widget.NewSelect(edgeOptions, func(changed string) {}), newDotSelector(0, 5, 0))
						}
					}
				},
				(*window),
			)
		}),
		widget.NewEntry(),
		widget.NewLabel(""),
		container.NewGridWithColumns(4, skillPrompts...),
		widget.NewLabel(""),
		container.NewGridWithColumns(2, edgePrompts...),
		widget.NewLabel(""),
		container.New(
			layout.NewFormLayout(),
			widget.NewLabel("Details"),
			widget.NewEntry(),
			widget.NewLabel("Path Condition"),
			widget.NewEntry(),
		),
	}
	return selector
}
