package tui

import (
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/rivo/tview"
)

func makeConfigFlex(h *hunter.Hunter) *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("top"), 0, 1, false).
		AddItem(makeCurrentConfigFlex(h), 0, 3, false)
	return flex
}

func makeCurrentConfigFlex(h *hunter.Hunter) *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("path"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("rules"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("verbose"), 0, 1, false)
	return flex
}
