package templates

import (
	// Blank embed needed to satisfy golint.
	_ "embed"
)

// These variables contain the embedded string value of
// the template for the associated output format.
var (
	//go:embed simple.tmpl
	Simple string

	//go:embed html.tmpl
	HTML string

	//go:embed markdown.tmpl
	Markdown string

	//go:embed table.tmpl
	Table string

	//go:embed html-table.tmpl
	HTMLTable string
)
