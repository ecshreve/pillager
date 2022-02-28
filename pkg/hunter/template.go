package hunter

import (
	"html/template"
	"log"
	"os"

	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v7/scan"
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
func RenderTemplate(w *os.File, tpl string, f scan.Report) {
	t, err := template.New("custom").Parse(tpl)
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to parse template"))
	}

	if err := t.Execute(w, f); err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to use custom template"))
	}
}
