package hunter

import (
	"io"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/zricethezav/gitleaks/v8/config"
	"github.com/zricethezav/gitleaks/v8/detect"
	"github.com/zricethezav/gitleaks/v8/report"
)

// Hunter is the secret scanner.
type Hunter struct {
	*Config
}

// New creates an instance of the Hunter.
func New(opts ...ConfigOption) (*Hunter, error) {
	return &Hunter{
		Config: NewConfig(opts...),
	}, nil
}

// Hunt walks over the filesystem at the configured path, looking for sensitive information.
func (h *Hunter) Hunt() ([]report.Finding, error) {
	output := detect.Options{Verbose: h.Verbose, Redact: h.Redact}
	conf := config.Config{Allowlist: h.Gitleaks.Allowlist, Rules: h.Gitleaks.Rules}

	findings, err := detect.FromFiles(h.ScanPath, conf, output)
	if err != nil {
		return nil, errors.Wrap(err, "failed to detect from files")
	}

	log.Debug().Bool("verbose", h.Verbose).Msg("scanner created")

	return findings, nil
}

// Report prints out the Findings in the preferred output format.
func (h *Hunter) Report(w io.Writer, findings []report.Finding) error {
	return h.Reporter.Report(w, findings)
}
