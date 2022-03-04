// Package rules enables the parsing of Gitleaks rulesets.
package rules

import (
	_ "embed"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"

	"github.com/zricethezav/gitleaks/v8/config"
)

// errReadConfig is the custom error message used if an error is encountered
// reading the gitleaks config.
const errReadConfig = "Failed to read gitleaks config"

// Default gitleaks configs, initialized at compile time via go:embed.
var (
	//go:embed rules_simple.toml
	RulesDefault string

	//go:embed rules_strict.toml
	RulesStrict string
)

// Loader represents a gitleaks config loader.
type Loader struct {
	loader config.ViperConfig
}

// LoaderOption sets a parameter for the gitleaks config loader.
type LoaderOption func(*Loader)

// NewLoader creates a Loader instance.
func NewLoader(opts ...LoaderOption) *Loader {
	var loader Loader
	if _, err := toml.Decode(RulesDefault, &loader.loader); err != nil {
		log.Fatal().Err(err).Msg(errReadConfig)
	}

	for _, opt := range opts {
		opt(&loader)
	}

	return &loader
}

// Load parses the gitleaks configuration.
func (l *Loader) Load() config.Config {
	config, err := l.loader.Translate()
	if err != nil {
		log.Fatal().Err(err).Msg(errReadConfig)
	}

	return config
}

// WithStrict decodes the default strict ruleset.
func WithStrict() LoaderOption {
	return func(l *Loader) {
		if _, err := toml.Decode(RulesStrict, &l.loader); err != nil {
			log.Fatal().Err(err).Msg(errReadConfig)
		}
	}
}

// FromFile decodes a gitleaks config from a local file.
func FromFile(file string) LoaderOption {
	return func(l *Loader) {
		if _, err := toml.DecodeFile(file, &l.loader); err != nil {
			log.Fatal().Err(err).Msg(errReadConfig)
		}
	}
}
