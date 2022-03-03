package tui

import "github.com/rivo/tview"

func makeOutputFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("output").SetTextAlign(0), 0, 1, false)
	return flex
}
