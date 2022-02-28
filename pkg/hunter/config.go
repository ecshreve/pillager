package hunter

import (
	"fmt"

	"github.com/brittonhayes/pillager/internal/validate"
	"github.com/brittonhayes/pillager/pkg/rules"
	"github.com/spf13/afero"
	gitleaks "github.com/zricethezav/gitleaks/v7/config"
)

// Config holds all the configurable parameters for a Hunter.
type Config struct {
	System   afero.Fs
	BasePath string
	Verbose  bool
	Workers  int
	Gitleaks gitleaks.Config
	Format   Format
	Template string
}

var _ Configer = &Config{}

// The Configer interface defines the available methods for instances of the
// Config type.
type Configer interface {
	Default() *Config
	Validate() (err error)
}

// NewConfig validates the given path and returns a new Config.
func NewConfig(fs afero.Fs, path string, verbose bool, gitleaks gitleaks.Config, format Format, template string, workers int) *Config {
	p := validate.New().Path(fs, path)
	return &Config{
		System:   fs,
		BasePath: p,
		Verbose:  verbose,
		Gitleaks: gitleaks,
		Format:   format,
		Template: template,
		Workers:  workers,
	}
}

// Default returns a Config with default values for the Hunter.
func (c *Config) Default() *Config {
	fs := afero.NewOsFs()
	v := validate.New()
	return &Config{
		System:   fs,
		BasePath: v.Path(fs, "."),
		Verbose:  false,
		Gitleaks: rules.Load(""),
		Format:   JSONFormat,
	}
}

// Validate returns an error if the given Config does not have the System
// or Rules fields populated.
func (c *Config) Validate() (err error) {
	if c.System == nil {
		err = fmt.Errorf("missing filesystem in Hunter Config")
	}

	if c.Gitleaks.Rules == nil {
		err = fmt.Errorf("no gitleaks config provided")
	}

	return
}
