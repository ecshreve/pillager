package hunter

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v8/report"
	"gopkg.in/yaml.v2"
)

type Leak report.Finding

// Report includes data related to the results of a Gitleaks Detect.
type Report struct {
	Leaks        []Leak
	SimpleString string
}

func (h *Hunter) BuildReport(findings []report.Finding) error {
	var leaks []Leak
	for _, f := range findings {
		leaks = append(leaks, Leak(f))
	}
	rep := &Report{
		Leaks: leaks,
	}

	if h.Config.Format == config.JSONFormat || h.Config.Format == config.YAMLFormat {
		simpleString, err := stringOutputHelper(h.Config.Format, leaks)
		if err != nil {
			return oops.Wrapf(err, "failed to build simple string for report")
		}
		rep.SimpleString = simpleString
	}

	h.Report = rep
	return nil
}

// Announce handles outputting a Hunter's Report.
func (h *Hunter) Announce() error {
	if !h.Config.Verbose {
		return nil
	}

	var buf bytes.Buffer
	if err := h.Config.Template.Execute(&buf, h.Report); err != nil {
		return oops.Wrapf(err, "failed to use custom template")
	}

	fmt.Printf("--- Results ---\n---\n")
	fmt.Println(buf.String())

	return nil
}

func stringOutputHelper(f config.Format, leaks []Leak) (string, error) {
	var (
		buf []byte
		err error
	)

	if f == config.JSONFormat {
		buf, err = json.Marshal(&leaks)
	} else if f == config.YAMLFormat {
		buf, err = yaml.Marshal(&leaks)
	} else {
		return "", oops.Errorf("invalid format for string output")
	}

	// Encountered an error unmarshalling data.
	if err != nil {
		return "", oops.Wrapf(err, "failed to unmarshal format: %s", f.String())
	}

	return string(buf), nil
}
