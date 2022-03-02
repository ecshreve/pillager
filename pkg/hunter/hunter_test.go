package hunter_test

import (
	"log"
	"testing"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/brittonhayes/pillager/templates"
	"github.com/samsarahq/go/oops"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHunter(t *testing.T) {
	// Basic Hunter creation works.
	hValid, err := hunter.NewHunter(ValidConfigForTest)
	assert.NoError(t, err)
	assert.NotNil(t, hValid)

	// Attempting to create a Hunter with invalid Config fails.
	hInvalid, err := hunter.NewHunter(InvalidConfigForTest)
	assert.Error(t, err)
	assert.Nil(t, hInvalid)
}

func TestHunt(t *testing.T) {
	// Basic Hunter for test.
	hunterForTest, err := hunter.NewHunter(ValidConfigForTest)
	require.NoError(t, err)
	require.NotNil(t, hunterForTest)

	// Hunt with valid Hunter works.
	err = hunterForTest.Hunt()
	assert.NoError(t, err)

	// Hunt with invalid Config returns error.
	hunterForTest.Config.BasePath = "asdfasdf"
	err = hunterForTest.Hunt()
	assert.Error(t, err)
	hunterForTest.Config.BasePath = "."
}

// This is an example of how to run a scan on a single file to look for
// email addresses.
func ExampleHunter_Hunt_simple() {
	env, err := HuntTestEnvHelper("./testdata/email.toml", "example@email.com")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}

	cfg, _ := NewTestConfig(env.TestFilePath, config.JSONFormat, templates.JSON)
	h, _ := hunter.NewHunter(*cfg)

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
	// --- Results ---
	// ---
	// [{"Description":"Email","StartLine":2,"EndLine":2,"StartColumn":1,"EndColumn":17,"Match":"example@email.com","Secret":"example@email.com","File":"./testdata/email.toml","Commit":"","Entropy":0,"Author":"","Email":"","Date":"","Message":"","Tags":["email"],"RuleID":""}]
	//
}

// This method also accepts custom output format configuration
// using go template/html. So if you don't like yaml or json, you can
// format to your heart's content.
func ExampleHunter_Hunt_template() {
	env, err := HuntTestEnvHelper("./testdata/package.toml", "https://github.com/brittonhayes/pillager")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "creating test env"))
	}

	cfg, _ := NewTestConfig(env.TestFilePath, config.CustomFormat, config.DefaultTemplate)
	h, _ := hunter.NewHunter(*cfg)

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
	// --- Results ---
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

	config, _ := NewTestConfig(env.TestFileContent, config.HTMLFormat, templates.HTML)
	h, _ := hunter.NewHunter(*config)

	if err = h.Hunt(); err != nil {
		log.Fatalln(oops.Wrapf(err, "failure to Hunt"))
	}
}
