package validate

import (
	"log"
	"os"

	"github.com/spf13/afero"
)

// Validation repreesents a Valid path.
type Validation struct{}

var _ Validator = &Validation{}

// The Validator interface defines available methods for instances of the
// Validation type.
type Validator interface {
	Path(fs afero.Fs, path string) string
}

// New returns a new Validation.
func New() *Validation {
	return &Validation{}
}

// Path checks if a filepath exists and returns it if so, otherwise it
// returns the default path.
func (v *Validation) Path(fs afero.Fs, path string) string {
	ok, err := afero.Exists(fs, path)
	if err != nil {
		log.Printf("ERROR: %s", err.Error())
		os.Exit(1)
	}

	if ok {
		return path
	}

	log.Fatal("no valid path provided")
	return "."
}
