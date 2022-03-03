package tui

import (
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/rivo/tview"
)

func makeConfigFlex(h *hunter.Hunter) *tview.Flex {
	aboutF := makeAboutFlex()
	aboutF.SetBorder(true).SetBorderPadding(1, 1, 2, 1)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(aboutF, 0, 1, false).
		AddItem(makeCurrentConfigFlex(h.Config), 0, 2, false)
	return flex
}

func makeAboutFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(ConfigBanner).SetTextAlign(0), 0, 1, false)
	return flex
}

func makeCurrentConfigFlex(c *hunter.Config) *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("path"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("rules"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("verbose"), 0, 1, false)
	return flex
}
