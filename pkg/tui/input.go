package tui

import "github.com/rivo/tview"

func makeInputFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("section1"), 0, 1, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("section2"), 0, 1, true).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("section3"), 0, 1, true)
	return flex
}
