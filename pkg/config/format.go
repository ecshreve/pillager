package config

import (
	"strings"

	"github.com/brittonhayes/pillager/templates"
)

// Format represents a possible output format for a Report.
type Format int

// All possible output formats for a Report.
const (
	JSONFormat Format = iota + 1
	YAMLFormat
	TableFormat
	HTMLFormat
	HTMLTableFormat
	MarkdownFormat
	CustomFormat
)

// FormatToTemplate is a Map of Format to the default output template string.
var FormatToTemplate = map[Format]string{
	JSONFormat:      templates.JSON,
	YAMLFormat:      templates.YAML,
	TableFormat:     templates.Table,
	HTMLFormat:      templates.HTML,
	HTMLTableFormat: templates.HTMLTable,
	MarkdownFormat:  templates.Markdown,
	CustomFormat:    templates.Simple,
}

// IsValid is a helper method to check the Format is one of the valid values.
func (f Format) IsValid() bool {
	switch f {
	case JSONFormat, YAMLFormat, TableFormat, HTMLFormat, HTMLTableFormat, MarkdownFormat, CustomFormat:
		return true
	}
	return false
}

// String implements the stringer interface for the Format type.
func (f Format) String() string {
	return [...]string{"", "json", "yaml", "table", "html", "html-table", "markdown", "custom"}[f]
}

// StringToFormat takes in a string representing the preferred output format,
// and returns the associated Format enum value.
func StringToFormat(s string) Format {
	switch strings.ToLower(s) {
	case "yaml":
		return YAMLFormat
	case "table":
		return TableFormat
	case "html":
		return HTMLFormat
	case "html-table":
		return HTMLTableFormat
	case "markdown":
		return MarkdownFormat
	case "custom":
		return CustomFormat
	default:
		return JSONFormat
	}
}