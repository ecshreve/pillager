package tui

import "github.com/rivo/tview"

func makeOutputFlex() *tview.Flex {
	return tview.NewFlex().
		AddItem(tview.NewTextView().
			SetText("Pillage filesystems for loot.").
			SetDynamicColors(true), 0, 1, false)
}
