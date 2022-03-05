package main

import "github.com/jroimartin/gocui"

type position struct {
	prc    float32
	margin int
}

func (p position) getCoordinate(max int) int {
	// value = prc * MAX + abs
	return int(p.prc*float32(max)) - p.margin
}

type viewPosition struct {
	x0, y0, x1, y1 position
}

func (vp viewPosition) getCoordinates(maxX, maxY int) (int, int, int, int) {
	var x0 = vp.x0.getCoordinate(maxX)
	var y0 = vp.y0.getCoordinate(maxY)
	var x1 = vp.x1.getCoordinate(maxX)
	var y1 = vp.y1.getCoordinate(maxY)
	return x0, y0, x1, y1
}

// All Views.
const (
	AboutView     = "about"
	OutputView    = "output"
	StatusView    = "status"
	ScanPathView  = "scan-path"
	RulesPathView = "rules-path"
	FormatView    = "format"
	TemplateView  = "template"
	WorkersView   = "workers"
	VerboseView   = "verbose"
	RedactView    = "redact"
)

var viewPositions = map[string]viewPosition{
	AboutView: {
		position{0.0, 0},
		position{0.0, 0},
		position{0.3, 2},
		position{0.3, 1},
	},
	ScanPathView: {
		position{0.0, 0},
		position{0.3, 0},
		position{0.3, 2},
		position{0.5, 1},
	},
	RulesPathView: {
		position{0.0, 0},
		position{0.4, 0},
		position{0.3, 2},
		position{0.6, 1},
	},
	FormatView: {
		position{0.0, 0},
		position{0.5, 1},
		position{0.3, 2},
		position{0.7, 1},
	},
	TemplateView: {
		position{0.0, 0},
		position{0.6, 0},
		position{0.3, 2},
		position{0.8, 1},
	},
	WorkersView: {
		position{0.0, 0},
		position{0.8, 0},
		position{0.1, 2},
		position{0.9, 1},
	},
	VerboseView: {
		position{0.1, 0},
		position{0.8, 0},
		position{0.2, 2},
		position{0.9, 1},
	},
	RedactView: {
		position{0.2, 0},
		position{0.8, 0},
		position{0.3, 2},
		position{0.9, 1},
	},
	OutputView: {
		position{0.3, 0},
		position{0.0, 0},
		position{1.0, 1},
		position{0.9, 1},
	},
	StatusView: {
		position{0.0, 0},
		position{0.9, 0},
		position{1.0, 1},
		position{1.0, 1},
	},
}

type ViewProps struct {
	title    string
	text     string
	frame    bool
	editable bool
	wrap     bool
	editor   gocui.Editor
}

func NewViewProps(title, text string, frame, editable, wrap bool, editor gocui.Editor) *ViewProps {
	return &ViewProps{
		title,
		text,
		frame,
		editable,
		wrap,
		editor,
	}
}
