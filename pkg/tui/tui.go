package tui

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var (
	app        *tview.Application
	configFlex *tview.Flex
	outputFlex *tview.Flex
	inputFlex  *tview.Flex
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

	// Initialize the flex boxes for our content.
	configFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	outputFlex = tview.NewFlex().SetDirection(tview.FlexColumn)
	inputFlex = tview.NewFlex().SetDirection(tview.FlexColumn)

	configFlex.SetBorder(true).SetTitle(" config ").SetBorderPadding(1, 1, 2, 1)
	outputFlex.SetBorder(true).SetTitle(" output ").SetBorderPadding(1, 1, 1, 1)
	inputFlex.SetBorder(true)

	configFlex.AddItem(makeConfigFlex(h), 0, 1, false)
	outputFlex.AddItem(makeOutputFlex(), 0, 1, false)
	inputFlex.AddItem(makeInputFlex(), 0, 1, false)

	// Attach our flex boxes to the outer flex container.
	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(configFlex, 10, 0, false).
			AddItem(outputFlex, 0, 3, false).
			AddItem(inputFlex, 0, 1, false), 0, 5, false)
	flex.SetBackgroundColor(tview.Styles.PrimitiveBackgroundColor)

	if err := app.SetRoot(flex, true).SetFocus(inputFlex).Run(); err != nil {
		panic(err)
	}
}
