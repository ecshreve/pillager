package hunter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"text/template"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v8/report"
	"gopkg.in/yaml.v2"
)

type Report struct {
	Leaks []report.Finding
}

// A Hound performs file inspection and collects the results.
type Hound struct {
	OutputFormat   config.Format
	CustomTemplate *string
	Findings       *Report `json:"findings"`
}

// NewHound creates an instance of the Hound type from the given Config.
func NewHound(f config.Format, t *string) *Hound {
	if f == config.CustomFormat && t == nil {
		log.Fatalln(oops.Errorf("invalid parameters for creating Hound"))
	}

	return &Hound{f, t, &Report{}}
}

// Howl prints out the findings from the Hound in the configured output format.
func (h *Hound) Howl() {
	fmt.Println("\n---\nHooooowl -- 🐕\n---")
	out := BuildOutputString(h.OutputFormat, h.CustomTemplate, *h.Findings)
	fmt.Println(out)
}

// RenderTemplate renders a Hound finding in a custom go template format to
// the provided writer.
func RenderTemplate(w io.Writer, tpl string, r Report) {
	t, err := template.New("custom").Parse(tpl)
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to parse template"))
	}

	if err := t.Execute(w, r); err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to use custom template"))
	}
}

// BuildOutputString returns the string result of applying the given template and
// format to the data in the report.
func BuildOutputString(f config.Format, tmp *string, rep Report) string {
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
