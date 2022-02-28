package config

import "strings"

// Format represents a possible output format for a Hound's findings.
type Format int

// All possible output formats for a Hound.
const (
	JSONFormat Format = iota + 1
	YAMLFormat
	TableFormat
	HTMLFormat
	HTMLTableFormat
	MarkdownFormat
	CustomFormat
)

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
