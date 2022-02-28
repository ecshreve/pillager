package hunter

import (
	"log"

	"github.com/samsarahq/go/oops"
	"github.com/zricethezav/gitleaks/v7/config"
	"github.com/zricethezav/gitleaks/v7/options"
	"github.com/zricethezav/gitleaks/v7/scan"
)

// Hunter holds the required fields to implement the Hunting interface
// and utilize the hunter package.
type Hunter struct {
	Config *Config
	Hound  *Hound
}

// NewHunter creates an instance of the Hunter type from the given Config.
func NewHunter(c *Config) *Hunter {
	if c == nil {
		conf := DefaultConfig()
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

	opt := options.Options{
		Path:    h.Config.BasePath,
		Verbose: h.Config.Verbose,
		Threads: h.Config.Workers,
	}
	conf := config.Config{
		Allowlist: h.Config.Gitleaks.Allowlist,
		Rules:     h.Config.Gitleaks.Rules,
	}

	scanner := scan.NewNoGitScanner(opt, conf)
	if scanner == nil {
		return oops.Errorf("unable to create scanner")
	}

	report, err := scanner.Scan()
	if err != nil {
		return oops.Wrapf(err, "unable to scan")
	}

	h.Hound.Findings = &report
	if opt.Verbose {
		h.Hound.Howl()
	}

	return nil
}
