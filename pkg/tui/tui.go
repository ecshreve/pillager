package tui

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app *tview.Application
)

// customCliTheme holds the color theme for the tview TUI.
var customCliTheme = tview.Theme{
	PrimitiveBackgroundColor:    tcell.Color(272727),
	ContrastBackgroundColor:     tcell.Color(448488),
	MoreContrastBackgroundColor: tcell.ColorGreen,
	BorderColor:                 tcell.ColorWhite,
	TitleColor:                  tcell.ColorWhite,
	GraphicsColor:               tcell.ColorWhite,
	PrimaryTextColor:            tcell.ColorWhite,
	SecondaryTextColor:          tcell.ColorYellow,
	TertiaryTextColor:           tcell.ColorGreen,
	InverseTextColor:            tcell.Color(448488),
	ContrastSecondaryTextColor:  tcell.ColorDarkCyan,
}

func StartTUI() {
	app = tview.NewApplication()
	tview.Styles = customCliTheme

	h, err := hunter.New()
	if err != nil {
		log.Fatalln(err)
	}

	configFlex := makeConfigFlex(h)
	contentFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(configFlex, 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("right"), 0, 3, false)

	inputFlex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("bottom"), 0, 1, true)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(contentFlex, 0, 4, false).
		AddItem(inputFlex, 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
