package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v8/config"
)

// Cfg holds all the configurable parameters for a Hunter.
type Cfg struct {
	BasePath string
	Verbose  bool
	Workers  int
	Format   Format
	Template string

	RulesPath string
	Rules     *gitleaks.Config
}

// NewCfg validates the given path and returns a new Config.
func NewCfg(path string, verbose bool, format Format, template string, workers int, rulesPath string, rules *gitleaks.Config) *Cfg {
	if rules == nil {
		parsedRules, err := ParseRules(rulesPath)
		if err != nil {
			log.Fatalln(oops.Wrapf(err, "parsing default pillager config file"))
		}
		rules = parsedRules
	}

	return &Cfg{
		BasePath:  path,
		Verbose:   verbose,
		Format:    format,
		Template:  template,
		Workers:   workers,
		RulesPath: rulesPath,
		Rules:     rules,
	}
}

// DefaultCfg returns a Cfg with default values for the Hunter.
func DefaultCfg() *Cfg {
	return NewCfg("", false, JSONFormat, "", 1, "", nil)
}

// Validate returns an error if the given Config does not have Rules defined.
func (c *Cfg) Validate() error {
	// If no file or directory exists at the given BasePath then set
	// it to the default value
	if _, err := os.Stat(c.BasePath); err != nil {
		c.BasePath = ""
	}

	if c.Rules.Rules == nil {
		return oops.Errorf("no gitleaks rules provided")
	}

	return nil
}

// ParseRules loads the rules defined in the rules.toml file
// into a list of gitleaks rules.
func ParseRules(filepath string) (*gitleaks.Config, error) {
	var loader gitleaks.ViperConfig

	if filepath != "" {
		if _, err := toml.DecodeFile(filepath, &loader); err != nil {
			return nil, oops.Wrapf(err, "failed to load config TOML data from file")
		}
	} else {
		if _, err := toml.Decode(DefaultRules, &loader); err != nil {
			return nil, oops.Wrapf(err, "failed to load default config TOML data")
		}
	}

	config, err := loader.Translate()
	if err != nil {
		return nil, oops.Wrapf(err, "failed to parse toml data to gitleaks config")
	}

	return &config, nil
}
