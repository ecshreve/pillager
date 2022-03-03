package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func makeInputFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewInputField().
			SetLabel("Enter a number: ").
			SetFieldWidth(0).
			SetDoneFunc(func(key tcell.Key) {
				app.Stop()
			}), 0, 1, true)

	return flex
}
