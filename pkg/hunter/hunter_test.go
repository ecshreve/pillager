package hunter_test

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
)

// This is an example of how to run a scan on a single file to look for
// email addresses.
func ExampleHunter_Hunt_simple() {
	env, err := HuntTestEnvHelper("./testdata/email.toml", "example@email.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}

	config := config.NewCfg("./testdata/email.toml", true, config.JSONFormat, config.DefaultTemplate, 1, "", env.Rules)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}

	// Output:
	// {
	// 	"Description": "Email",
	// 	"StartLine": 2,
	// 	"EndLine": 2,
	// 	"StartColumn": 1,
	// 	"EndColumn": 17,
	// 	"Match": "example@email.com",
	// 	"Secret": "example@email.com",
	//	"File": "./testdata/email.toml",
	// 	"Commit": "",
	// 	"Entropy": 0,
	// 	"Author": "",
	// 	"Email": "",
	// 	"Date": "",
	// 	"Message": "",
	// 	"Tags": [
	// 		"email"
	// 	],
	// 	"RuleID": ""
	// }
	//
	// ---
	// Hooooowl -- 🐕
	// ---
	// [{"Description":"Email","StartLine":2,"EndLine":2,"StartColumn":1,"EndColumn":17,"Match":"example@email.com","Secret":"example@email.com","File":"./testdata/email.toml","Commit":"","Entropy":0,"Author":"","Email":"","Date":"","Message":"","Tags":["email"],"RuleID":""}]
}

// This method also accepts custom output format configuration
// using go template/html. So if you don't like yaml or json, you can
// format to your heart's content.
func ExampleHunter_Hunt_template() {
	env, err := HuntTestEnvHelper("./testdata/package.toml", "https://github.com/brittonhayes/pillager")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}

	config := config.NewCfg("./testdata/package.toml", true, config.CustomFormat, config.DefaultTemplate, 1, "", env.Rules)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}

	// Output:
	// {
	// 	"Description": "Github",
	// 	"StartLine": 2,
	// 	"EndLine": 2,
	// 	"StartColumn": 1,
	// 	"EndColumn": 40,
	// 	"Match": "https://github.com/brittonhayes/pillager",
	// 	"Secret": "https://github.com/brittonhayes/pillager",
	//	"File": "./testdata/package.toml",
	// 	"Commit": "",
	// 	"Entropy": 0,
	// 	"Author": "",
	// 	"Email": "",
	// 	"Date": "",
	// 	"Message": "",
	// 	"Tags": [
	// 		"github"
	// 	],
	// 	"RuleID": ""
	// }
	//
	// ---
	// Hooooowl -- 🐕
	// ---
	// Line: 2
	// File: ./testdata/package.toml
	// Offender: https://github.com/brittonhayes/pillager
}

// Hunter will also look personally identifiable info in TOML files and
// format the output as HTML.
func ExampleHunter_Hunt_toml() {
	env, err := HuntTestEnvHelper("./testdata/email.toml", "fakeperson@example.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}

	config := config.NewCfg(env.TestFilePath, true, config.HTMLFormat, templates.HTML, 1, "", env.Rules)
	h := hunter.NewHunter(config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}
}
