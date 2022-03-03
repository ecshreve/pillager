package tui

import (
	"fmt"
	"path/filepath"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func makeConfigFlex(c *hunter.Config) *tview.Flex {
	aboutFlex := makeAboutFlex()
	currentFlex := makeCurentFlex(c)
	aboutFlex.SetBorder(true).SetBorderColor(tcell.ColorLightBlue).SetBorderPadding(1, 1, 1, 1)
	currentFlex.SetBorder(true).SetTitle(" Current Config ").SetBorderPadding(0, 0, 1, 1)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(aboutFlex, 0, 1, false).
		AddItem(currentFlex, 0, 2, false)
	return flex
}

func makeAboutFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(ConfigBanner).SetTextAlign(1), 0, 1, false)
	return flex
}

func makeCurentFlex(c *hunter.Config) *tview.Flex {
	absScanPath, err := filepath.Abs(c.ScanPath)
	if err != nil {
		absScanPath = c.ScanPath
	}

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(makeConfigItem("ScanPath", absScanPath), 0, 1, false).
		AddItem(makeConfigItem("Rules", c.Gitleaks.Description), 0, 1, false).
		AddItem(makeConfigItem("Format", fmt.Sprintf("%T", c.Reporter)[7:]), 3, 0, false).
		AddItem(makeConfigItem("NumWorkers", fmt.Sprintf("%d", c.Workers)), 3, 0, false).
		AddItem(makeConfigItem("Verbose", fmt.Sprintf("%v", c.Verbose)), 3, 0, false).
		AddItem(makeConfigItem("Redact", fmt.Sprintf("%v", c.Redact)), 3, 0, false)

	return flex
}

func makeConfigItem(title, text string) *tview.TextView {
	t := tview.NewTextView().SetText(text).SetTextColor(tcell.ColorLightGreen).SetWrap(true)
	t.SetBorder(true).SetBorderPadding(0, 0, 1, 1).
		SetTitle(fmt.Sprintf(" %s ", title)).SetTitleAlign(0).SetTitleColor(tcell.ColorGray)
	return t
}
