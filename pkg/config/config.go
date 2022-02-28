package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v7/config"
)

// Cfg holds all the configurable parameters for a Hunter.
type Cfg struct {
	BasePath string
	Verbose  bool
	Workers  int
	Gitleaks *gitleaks.Config
	Format   Format
	Template string
}

// NewCfg validates the given path and returns a new Config.
func NewCfg(path string, verbose bool, gitleaks *gitleaks.Config, format Format, template string, workers int) *Cfg {
	return &Cfg{
		BasePath: path,
		Verbose:  verbose,
		Gitleaks: gitleaks,
		Format:   format,
		Template: template,
		Workers:  workers,
	}
}

// DefaultCfg returns a Cfg with default values for the Hunter.
func DefaultCfg() *Cfg {
	gitleaks, err := ParsePillagerConfigFile("")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "parsing default pillager config file"))
	}

	return &Cfg{
		BasePath: "",
		Verbose:  false,
		Gitleaks: gitleaks,
		Format:   JSONFormat,
	}
}

// Validate returns an error if the given Config does not have the System
// or Rules fields populated.
func (c *Cfg) Validate() error {
	// If no file or directory exists at the given BasePath then set
	// it to the default value
	if _, err := os.Stat(c.BasePath); err != nil {
		c.BasePath = ""
	}

	if c.Gitleaks.Rules == nil {
		return oops.Errorf("no gitleaks rules provided")
	}

	return nil
}

// ParsePillagerConfigFile loads the rules defined in the config file
// into a list of gitleaks rules.
func ParsePillagerConfigFile(filepath string) (*gitleaks.Config, error) {
	var loader gitleaks.TomlLoader

	if filepath != "" {
		if _, err := toml.DecodeFile(filepath, &loader); err != nil {
			return nil, oops.Wrapf(err, "failed to load config TOML data from file")
		}
	} else {
		if _, err := toml.Decode(DefaultPillagerConfig, &loader); err != nil {
			return nil, oops.Wrapf(err, "failed to load default config TOML data")
		}
	}

	config, err := loader.Parse()
	if err != nil {
		return nil, oops.Wrapf(err, "failed to parse toml data to gitleaks config")
	}

	return &config, nil
}
