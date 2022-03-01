package hunter

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v8/detect"
)

// Hunter holds configuration and reference to a Hound.
type Hunter struct {
	Config *config.Cfg
	Hound  *Hound
}

// NewHunter creates an instance of the Hunter type from the given Config.
func NewHunter(c *config.Cfg) *Hunter {
	if c == nil {
		conf := config.DefaultCfg()
		return &Hunter{conf, NewHound(conf.Format, &conf.Template)}
	}

	if err := c.Validate(); err != nil {
		log.Fatalln(oops.Wrapf(err, "invalid Config"))
	}

	return &Hunter{c, NewHound(c.Format, &c.Template)}
}

// Hunt walks over the filesystem at the configured path, looking for
// sensitive information.
func (h *Hunter) Hunt() error {
	if h.Hound == nil {
		h.Hound = NewHound(h.Config.Format, &h.Config.Template)
	}

	opt := detect.Options{
		Verbose: h.Config.Verbose,
	}

	findings, err := detect.FromFiles(h.Config.BasePath, *h.Config.Rules, opt)
	if err != nil {
		return oops.Wrapf(err, "detecting findings")
	}

	h.Hound.Findings = &Report{Leaks: findings}

	if h.Config.Verbose {
		h.Hound.Howl()
	}

	return nil
}
