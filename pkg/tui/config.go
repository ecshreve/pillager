package tui

import (
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func makeConfigFlex(c *hunter.Config) *tview.Flex {
	aboutFlex := makeAboutFlex()
	currentFlex := makeCurentFlex(c)
	aboutFlex.SetBorder(true).SetTitle(" about ").SetBorderPadding(1, 1, 1, 1)
	currentFlex.SetBorder(true).SetTitle(" current ").SetBorderPadding(1, 1, 1, 1)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(aboutFlex, 0, 1, false).
		AddItem(currentFlex, 0, 2, true)
	return flex
}

func makeAboutFlex() *tview.Flex {
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText(ConfigBanner).SetTextAlign(0), 0, 1, false)
	return flex
}

func makeCurentFlex(c *hunter.Config) *tview.Flex {
	form := tview.NewForm().
		AddInputField("ScanPath", c.ScanPath, 0, nil, nil).
		AddInputField("RulesPath", "", 20, nil, nil).
		AddInputField("Template", "", 20, nil, nil).
		AddCheckbox("Verbose", true, nil).
		AddCheckbox("Redact", false, nil).
		AddInputField("Workers", "", 0, nil, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	form.SetFieldBackgroundColor(tcell.ColorBlue)
	form.SetBorder(true).SetTitle("Enter some data").SetTitleAlign(tview.AlignLeft)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(form, 0, 1, true)
	return flex
}
