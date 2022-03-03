package tui

import (
	"os"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app *tview.Application

	contentFlex *tview.Flex
	inputFlex   *tview.Flex

	inputField *tview.InputField
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

	c := hunter.NewConfig()

	contentFlex = makeContentFlex(c)
	inputFlex = makeInputFlex()
	contentFlex.SetBorder(true).SetTitle(" content ").SetBorderPadding(1, 1, 1, 1)
	inputFlex.SetBorder(true).SetTitle(" input ").SetBorderPadding(1, 1, 1, 1)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(contentFlex, 0, 3, false).
		AddItem(inputFlex, 0, 1, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'c':
			app.SetFocus(contentFlex)
		case 'q':
			app.Stop()
			os.Exit(0)
		}

		return event
	})

	if err := app.SetRoot(flex, true).SetFocus(inputFlex).Run(); err != nil {
		panic(err)
	}
}
