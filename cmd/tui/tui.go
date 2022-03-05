package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/jroimartin/gocui"
)

type App struct {
	hunter       *hunter.Hunter
	viewIndex    int
	historyIndex int
	currentPopup string
	history      []string
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	h, _ := hunter.New()
	app := &App{hunter: h}
	g.SetManagerFunc(app.Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(*gocui.Gui, *gocui.View) error {
		v, err := g.SetCurrentView(OutputView)
		if err != nil {
			return err
		}

		r, err := h.Hunt()
		if err != nil {
			return err
		}

		a, err := json.MarshalIndent(r, "", "    ")
		if err != nil {
			return err
		}

		fmt.Fprint(v, string(a))

		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func (a *App) Layout(g *gocui.Gui) error {
	var views = []string{
		AboutView,
		OutputView,
		StatusView,
		ScanPathView,
		RulesPathView,
		FormatView,
		TemplateView,
		WorkersView,
		VerboseView,
		RedactView,
	}

	hunterConfigMap := HunterConfigToMap(a.hunter)

	maxX, maxY := g.Size()
	for _, view := range views {
		x0, y0, x1, y1 := viewPositions[view].getCoordinates(maxX, maxY)
		if v, err := g.SetView(view, x0, y0, x1, y1); err != nil {
			v.SelFgColor = gocui.ColorBlack
			v.SelBgColor = gocui.ColorGreen

			v.Title = " " + view + " "
			v.Clear()
			fmt.Fprintf(v, hunterConfigMap[view])
			if err != gocui.ErrUnknownView {
				return err
			}
		}
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
