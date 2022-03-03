package tui

import (
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/rivo/tview"
)

func makeContentFlex(c *hunter.Config) *tview.Flex {
	configFlex := makeConfigFlex(c)
	outputFlex := makeOutputFlex()
	configFlex.SetBorder(true).SetTitle(" config ").SetBorderPadding(1, 1, 1, 1)
	outputFlex.SetBorder(true).SetTitle(" output ").SetBorderPadding(1, 1, 1, 1)

	flex := tview.NewFlex().
		AddItem(configFlex, 0, 1, true).
		AddItem(outputFlex, 0, 2, false)
	return flex
}
