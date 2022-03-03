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

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(makeConfigFlex(h), 0, 1, false).
			AddItem(makeOutputFlex(), 0, 3, false), 0, 4, false).
		AddItem(makeInputFlex(), 0, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
		panic(err)
	}
}
