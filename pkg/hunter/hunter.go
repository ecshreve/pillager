package hunter

import (
	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v8/detect"
)

// Hunter holds configuration and reference to a Hound.
type Hunter struct {
	Config *config.Config
	Report *Report
}

// NewHunter creates an instance of the Hunter type from the given Config.
func NewHunter(c config.Config) (*Hunter, error) {
	if err := c.Validate(); err != nil {
		return nil, oops.Wrapf(err, "failed to create new Hunter with invalid config: %+v", c)
	}

	return &Hunter{
		Config: &c,
	}, nil
}

// Hunt walks over the filesystem at the configured path, looking for
// sensitive information. Checking each file in under the configured path
// is accomplished with Gitleaks Detect functionality.
func (h *Hunter) Hunt() error {
	findings, err := detect.FromFiles(
		h.Config.BasePath,
		*h.Config.Rules,
		detect.Options{
			Verbose: h.Config.Verbose,
		},
	)
	if err != nil {
		return oops.Wrapf(err, "unable to detect findings")
	}

	if err = h.BuildReport(findings); err != nil {
		return oops.Wrapf(err, "unable to build report")
	}

	h.Announce()

	return nil
}
