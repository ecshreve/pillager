package hunter_test

import (
	"log"
	"os"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v7/config"
)

// HunterTestEnv is a convenient mechanism to handle test environment
// setup and cleanup.
type HunterTestEnv struct {
	Gitleaks        *gitleaks.Config
	TestFileName    string
	TestFileContent string
}

func (e *HunterTestEnv) cleanup() {
	if err := os.Remove(e.TestFileName); err != nil {
		log.Println(oops.Wrapf(err, "removing temporary test files"))
	}
}

// huntTestEnvHelper is a convenient helper to create temporary files
// with for the tests and examples in this package.
func HuntTestEnvHelper(testFilePattern string, testFileContent string) (*HunterTestEnv, error) {
	gl, err := hunter.ParseRulesFromConfigFile("./testdata/pillager_test_config.toml")
	if err != nil {
		return nil, oops.Wrapf(err, "parsing config data for test env")
	}
	// Create test file to scan and write some data into it.
	f, err := os.CreateTemp("./testdata", testFilePattern)
	if err != nil {
		return nil, oops.Wrapf(err, "creating test file for test env")
	}
	defer f.Close()

	_, err = f.WriteString(testFileContent)
	if err != nil {
		return nil, oops.Wrapf(err, "writing test file content for test env")
	}

	return &HunterTestEnv{
		Gitleaks:        gl,
		TestFileName:    f.Name(),
		TestFileContent: testFileContent,
	}, nil
}
