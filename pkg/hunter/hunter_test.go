package hunter_test

import (
	"log"
	"os"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
	gitleaks "github.com/zricethezav/gitleaks/v7/config"
)

type hunterTestEnv struct {
	gitleaks        *gitleaks.Config
	testFileName    string
	testFileContent string
}

func (e *hunterTestEnv) cleanup() error {
	return os.Remove(e.testFileName)
}

// huntTestEnvHelper is a convenient helper to create temporary files
// with for the tests and examples in this package.
func huntTestEnvHelper(testFilePattern string, testFileContent string) (*hunterTestEnv, error) {
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

	return &hunterTestEnv{
		gitleaks:        gl,
		testFileName:    f.Name(),
		testFileContent: testFileContent,
	}, nil
}

// This is an example of how to run a scan on a single file to look for
// email addresses.
func ExampleHunter_Hunt_email() {
	env, err := huntTestEnvHelper("~.toml", "example@email.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := hunter.NewConfig(env.testFileName, true, env.gitleaks, hunter.JSONFormat, hunter.DefaultTemplate, 1)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}

	// Output:
	// {
	// 	"line": "example@email.com",
	// 	"lineNumber": 1,
	// 	"offender": "example@email.com",
	// 	"offenderEntropy": -1,
	// 	"commit": "",
	// 	"repo": "",
	// 	"repoURL": "",
	// 	"leakURL": "",
	// 	"rule": "Email",
	// 	"commitMessage": "",
	// 	"author": "",
	// 	"email": "",
	// 	"file": ".",
	// 	"date": "0001-01-01T00:00:00Z",
	// 	"tags": "email"
	// }
	//
	// ---
	// Hooooowl -- 🐕
	// ---
	// [{"line":"example@email.com","lineNumber":1,"offender":"example@email.com","offenderEntropy":-1,"commit":"","repo":"","repoURL":"","leakURL":"","rule":"Email","commitMessage":"","author":"","email":"","file":".","date":"0001-01-01T00:00:00Z","tags":"email"}]
	//
}

// This method also accepts custom output format configuration
// using go template/html. So if you don't like yaml or json, you can
// format to your heart's content.
func ExampleHunter_Hunt_custom_output() {
	env, err := huntTestEnvHelper("~.yaml", "https://github.com/brittonhayes/pillager")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := hunter.NewConfig(env.testFileName, true, env.gitleaks, hunter.CustomFormat, hunter.DefaultTemplate, 1)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}

	// Output:
	// {
	// 	"line": "https://github.com/brittonhayes/pillager",
	// 	"lineNumber": 1,
	// 	"offender": "https://github.com/brittonhayes/pillager",
	// 	"offenderEntropy": -1,
	// 	"commit": "",
	// 	"repo": "",
	// 	"repoURL": "",
	// 	"leakURL": "",
	// 	"rule": "Github",
	// 	"commitMessage": "",
	// 	"author": "",
	// 	"email": "",
	// 	"file": ".",
	// 	"date": "0001-01-01T00:00:00Z",
	// 	"tags": "github"
	// }
	//
	// ---
	// Hooooowl -- 🐕
	// ---
	// Line: 1
	// File: .
	// Offender: https://github.com/brittonhayes/pillager
}

// Hunter will also look personally identifiable info in TOML files and
// format the output as HTML.
func ExampleHunter_Hunt_toml() {
	env, err := huntTestEnvHelper("~.toml", "fakeperson@example.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := hunter.NewConfig(env.testFileName, true, env.gitleaks, hunter.HTMLFormat, templates.HTML, 1)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}
}
