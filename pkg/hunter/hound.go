package hunter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/templates"
	"github.com/ghodss/yaml"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v7/scan"
)

// A Hound performs file inspection and collects the results.
type Hound struct {
	OutputFormat   config.Format
	CustomTemplate *string
	Findings       *scan.Report `json:"findings"`
}

// NewHound creates an instance of the Hound type from the given Config.
func NewHound(f config.Format, t *string) *Hound {
	if f == config.CustomFormat && t == nil {
		log.Fatalln(oops.Errorf("invalid parameters for creating Hound"))
	}

	return &Hound{f, t, &scan.Report{}}
}

// Howl prints out the findings from the Hound in the configured output format.
func (h *Hound) Howl() {
	fmt.Println("\n---\nHooooowl -- 🐕\n---")
	findings := *h.Findings

	switch h.OutputFormat {
	case config.JSONFormat:
		b, err := json.Marshal(&findings.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from json"))
		}
		fmt.Println(string(b))
	case config.YAMLFormat:
		b, err := yaml.Marshal(&findings.Leaks)
		if err != nil {
			log.Fatal(oops.Wrapf(err, "unmarshal findings from yaml"))
		}
		fmt.Println(string(b))
	case config.HTMLFormat:
		RenderTemplate(os.Stdout, templates.HTML, findings)
	case config.HTMLTableFormat:
		RenderTemplate(os.Stdout, templates.HTMLTable, findings)
	case config.MarkdownFormat:
		RenderTemplate(os.Stdout, templates.Markdown, findings)
	case config.TableFormat:
		RenderTemplate(os.Stdout, templates.Table, findings)
	case config.CustomFormat:
		if h.CustomTemplate == nil {
			log.Fatalln(oops.Errorf("unable to execute Howl for CustomFormat with nil CustomTemplate"))
		}
		RenderTemplate(os.Stdout, *h.CustomTemplate, findings)
	default:
		RenderTemplate(os.Stdout, templates.Simple, findings)
	}
}
