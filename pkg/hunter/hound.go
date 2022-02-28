package hunter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/pillager/templates"
	"github.com/ghodss/yaml"
	"github.com/zricethezav/gitleaks/v7/scan"
)

// A Hound performs file inspection and collects the results.
type Hound struct {
	Config   *Config
	Findings scan.Report `json:"findings"`
}

var _ Hounder = &Hound{}

// The Hounder interface defines the available methods for instances of the
// Hound type.
type Hounder interface {
	Howl(findings scan.Report)
}

// NewHound creates an instance of the Hound type from the given Config.
func NewHound(c *Config) *Hound {
	if c == nil {
		var conf Config
		return &Hound{conf.Default(), scan.Report{}}
	}

	if c.System == nil {
		log.Fatal("Missing filesystem in Hunter Config")
	}

	return &Hound{c, scan.Report{}}
}

// Howl prints out the findings from the Hound in the configured output format.
func (h *Hound) Howl(findings scan.Report) {
	if h.Config.Template != "" {
		h.Config.Format = CustomFormat
	}
	switch h.Config.Format {
	case JSONFormat:
		b, err := json.Marshal(&findings.Leaks)
		if err != nil {
			log.Fatal("Failed to unmarshal findings")
		}
		fmt.Println(string(b))
	case YAMLFormat:
		b, err := yaml.Marshal(&findings.Leaks)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Println(string(b))
	case HTMLFormat:
		RenderTemplate(os.Stdout, templates.HTML, findings)
	case HTMLTableFormat:
		RenderTemplate(os.Stdout, templates.HTMLTable, findings)
	case MarkdownFormat:
		RenderTemplate(os.Stdout, templates.Markdown, findings)
	case TableFormat:
		RenderTemplate(os.Stdout, templates.Table, findings)
	case CustomFormat:
		RenderTemplate(os.Stdout, h.Config.Template, findings)
	default:
		RenderTemplate(os.Stdout, templates.Simple, findings)
	}
}
