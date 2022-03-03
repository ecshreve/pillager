package tui

import "github.com/rivo/tview"

func makeInputFlex() *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewTextView().
			SetText("Input").
			SetDynamicColors(true), 0, 1, false)
}
