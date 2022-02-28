package hunter

import (
	"fmt"
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/helpers"
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
	out := helpers.BuildOutputString(h.OutputFormat, h.CustomTemplate, *h.Findings)
	fmt.Println(out)
}
