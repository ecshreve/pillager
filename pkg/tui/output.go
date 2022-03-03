package tui

import "github.com/rivo/tview"

func makeOutputFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("right"), 0, 1, false)

	return flex
}
