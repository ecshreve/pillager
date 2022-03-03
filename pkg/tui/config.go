package tui

import (
	"fmt"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/rivo/tview"
)

func makeConfigFlex(h *hunter.Hunter) *tview.Flex {
	ff := tview.NewFlex().SetDirection(tview.FlexColumn)

	ff.AddItem(tview.NewTextView().SetText(ConfigBanner).SetDynamicColors(true), 0, 1, false)
	ff.AddItem(makeCurrentConfigFlex(h), 0, 1, false)

	return ff
}

func makeCurrentConfigFlex(h *hunter.Hunter) *tview.Flex {
	tst := fmt.Sprintf("%+v", h.Config)
	return tview.NewFlex().
		AddItem(tview.NewTextView().SetText(tst).SetDynamicColors(true), 0, 1, false)
}
