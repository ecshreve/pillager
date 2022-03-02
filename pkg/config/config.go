package config

import (
	"log"
	"os"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v8/config"
)

type ConfigParams struct {
	BasePath  string
	RulesPath string
	Format    Format
	Verbose   bool
	Workers   int
	Template  string
}

// Config holds parameters used by a Hunter.
type Config struct {
	BasePath string

	Format  Format
	Verbose bool
	Workers int

	Template *template.Template
	Rules    *gitleaks.Config
}

// NewConfig validates the given path and returns a new Config.
func NewConfig(p ConfigParams) (*Config, error) {
	parsedRules, err := ParseRules(p.RulesPath)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to parse rules file")
	}

	parsedTemplate, err := template.New("out-template").Parse(p.Template)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create parsed template")
	}

	c := &Config{
		BasePath: p.BasePath,
		Format:   p.Format,
		Verbose:  p.Verbose,
		Workers:  p.Workers,
		Template: parsedTemplate,
		Rules:    parsedRules,
	}

	if err := c.Validate(); err != nil {
		return nil, oops.Wrapf(err, "unable to validate config: %v from params: %v", c, p)
	}

	return c, nil
}

// DefaultConfig returns a Cfg with default values for the Hunter.
func DefaultConfig() *Config {
	defaultConfigParams := &ConfigParams{
		BasePath:  "",
		RulesPath: "",
		Verbose:   false,
		Format:    JSONFormat,
		Template:  "",
		Workers:   1,
	}

	defaultConfig, err := NewConfig(*defaultConfigParams)
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "failed to create default config"))
	}

	return defaultConfig
}

// Validate returns an error if the given Config doesn't have valid field values.
func (c *Config) Validate() error {
	if _, err := os.Stat(c.BasePath); err != nil {
		return oops.Errorf("path does not exist")
	}

	if !c.Format.IsValid() {
		return oops.Errorf("invalid format")
	}

	if c.Workers < 1 || c.Workers > 100 {
		return oops.Errorf("number of workers out of bounds")
	}

	if c.Rules == nil {
		return oops.Errorf("no gitleaks rules")
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
