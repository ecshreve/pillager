package hunter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/pillager/templates"
	"github.com/ghodss/yaml"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v7/scan"
)

// A Hound performs file inspection and collects the results.
type Hound struct {
	OutputFormat   Format
	CustomTemplate *string
	Findings       *scan.Report `json:"findings"`
}

// NewHound creates an instance of the Hound type from the given Config.
func NewHound(f Format, t *string) *Hound {
	if f == CustomFormat && t == nil {
		log.Fatalln(oops.Errorf("invalid parameters for creating Hound"))
	}

	return &Hound{f, t, &scan.Report{}}
}

// Howl prints out the findings from the Hound in the configured output format.
func (h *Hound) Howl() {
	fmt.Println("\n---\nHooooowl -- 🐕\n---")
	findings := *h.Findings

	switch h.OutputFormat {
	case JSONFormat:
		b, err := json.Marshal(&findings.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from json"))
		}
		fmt.Println(string(b))
	case YAMLFormat:
		b, err := yaml.Marshal(&findings.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from yaml"))
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
		if h.CustomTemplate == nil {
			log.Fatalln(oops.Errorf("unable to execute Howl for CustomFormat with nil CustomTemplate"))
		}
		RenderTemplate(os.Stdout, *h.CustomTemplate, findings)
	default:
		RenderTemplate(os.Stdout, templates.Simple, findings)
	}
}
