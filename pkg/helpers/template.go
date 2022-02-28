package helpers

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v7/scan"
	"gopkg.in/yaml.v2"
)

// DefaultTemplate is the base template used to format a Finding into the
// custom output format.
const DefaultTemplate = `{{ with . -}}
{{ range .Leaks -}}
Line: {{.LineNumber}}
File: {{ .File }}
Offender: {{ .Offender }}

{{ end -}}{{- end}}`

// RenderTemplate renders a Hound finding in a custom go template format to
// the provided writer.
func RenderTemplate(w io.Writer, tpl string, f scan.Report) {
	t, err := template.New("custom").Parse(tpl)
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to parse template"))
	}

	if err := t.Execute(w, f); err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to use custom template"))
	}
}

// BuildOutputString returns the string result of applying the given template and
// format to the data in the report.
func BuildOutputString(f config.Format, tmp *string, rep scan.Report) string {
	var buf bytes.Buffer

	switch f {
	case config.JSONFormat:
		b, err := json.Marshal(&rep.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from json"))
		}
		return string(b)
	case config.YAMLFormat:
		b, err := yaml.Marshal(&rep.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from yaml"))
		}
		return string(b)
	case config.HTMLFormat:
		RenderTemplate(&buf, templates.HTML, rep)
	case config.HTMLTableFormat:
		RenderTemplate(&buf, templates.HTMLTable, rep)
	case config.MarkdownFormat:
		RenderTemplate(&buf, templates.Markdown, rep)
	case config.TableFormat:
		RenderTemplate(&buf, templates.Table, rep)
	case config.CustomFormat:
		if tmp == nil {
			log.Fatalln(oops.Errorf("unable to build output string for CustomFormat with nil CustomTemplate"))
		}
		RenderTemplate(&buf, *tmp, rep)
	default:
		RenderTemplate(&buf, templates.Simple, rep)

	}

	return buf.String()
}
