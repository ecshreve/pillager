package hunter_test

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/helpers"
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
)

// This is an example of how to run a scan on a single file to look for
// email addresses.
func ExampleHunter_Hunt_simple() {
	env, err := HuntTestEnvHelper("~.toml", "example@email.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := config.NewCfg(env.TestFileName, true, env.Gitleaks, config.JSONFormat, helpers.DefaultTemplate, 1)
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
func ExampleHunter_Hunt_template() {
	env, err := HuntTestEnvHelper("~.yaml", "https://github.com/brittonhayes/pillager")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := config.NewCfg(env.TestFileName, true, env.Gitleaks, config.CustomFormat, helpers.DefaultTemplate, 1)
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
	env, err := HuntTestEnvHelper("~.toml", "fakeperson@example.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}
	defer env.cleanup()

	config := config.NewCfg(env.TestFileName, true, env.Gitleaks, config.HTMLFormat, templates.HTML, 1)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}
}
