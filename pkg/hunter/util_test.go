package hunter_test

import (
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v8/config"
)

// HunterTestEnv is a convenient mechanism to handle test environment
// setup and cleanup.
type HunterTestEnv struct {
	Rules           *gitleaks.Config
	TestFilePath    string
	TestFileContent string
}

// huntTestEnvHelper is a convenient helper to create temporary files
// with for the tests and examples in this package.
func HuntTestEnvHelper(testFilePath string, testFileContent string) (*HunterTestEnv, error) {
	r, err := GetRulesForTest()
	if err != nil {
		return nil, oops.Wrapf(err, "getting config data for test env")
	}

	return &HunterTestEnv{
		Rules:           r,
		TestFilePath:    testFilePath,
		TestFileContent: testFileContent,
	}, nil
}

// GetPillagerConfigForTest returns a basic gitleaks config for use in tests.
func GetRulesForTest() (*gitleaks.Config, error) {
	var loader gitleaks.ViperConfig
	if _, err := toml.Decode(config.RulesForTest, &loader); err != nil {
		return nil, oops.Wrapf(err, "failed to load default config toml data")
	}

	config, err := loader.Translate()
	if err != nil {
		return nil, oops.Wrapf(err, "failed to parse toml data to gitleaks config")
	}

	return &config, nil
}

func NewTestConfig(path string, f config.Format, t string) (*config.Config, error) {
	parsedRules, err := GetRulesForTest()
	if err != nil {
		return nil, oops.Wrapf(err, "failed to parse rules file")
	}

	parsedTemplate, err := template.New("out-template").Parse(t)
	if err != nil {
		return nil, oops.Wrapf(err, "failed to create parsed template")
	}

	return &config.Config{
		BasePath: path,
		Format:   f,
		Verbose:  true,
		Workers:  1,
		Template: parsedTemplate,
		Rules:    parsedRules,
	}, nil
}

var ValidConfigForTest = config.Config{
	BasePath: ".",
	Format:   config.JSONFormat,
	Verbose:  false,
	Workers:  1,
	Template: nil,
	Rules:    &gitleaks.Config{Rules: []*gitleaks.Rule{}},
}

var InvalidConfigForTest = config.Config{
	BasePath: "asdfasd",
	Format:   config.JSONFormat,
	Verbose:  false,
	Workers:  1,
	Template: nil,
	Rules:    &gitleaks.Config{Rules: []*gitleaks.Rule{}},
}
