package hunter_test

import (
	"log"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/brittonhayes/pillager/templates"
)

// Here is an example of utilizing the Announce function on a
// Report. The Announce method is the final method in the
// hunting process. It takes whatever has been found and
// outputs it for the user.
func ExampleHunter_Announce_table() {
	cfg, err := NewTestConfig(".", config.TableFormat, templates.Table)
	if err != nil {
		log.Fatalln(err)
	}

	h, err := hunter.NewHunter(*cfg)
	if err != nil {
		log.Fatalln(err)
	}

	h.Report = &hunter.Report{
		Leaks: []hunter.Leak{
			{
				Secret:    "person@email.com",
				StartLine: 16,
				Match:     "person@email.com",
				RuleID:    "Email Addresses",
				File:      "example.txt",
			},
			{
				Secret:    "fred@email.com",
				StartLine: 29,
				Match:     "fred@email.com",
				RuleID:    "Email Addresses",
				File:      "example2.txt",
			},
		},
	}

	err = h.Announce()
	if err != nil {
		log.Fatalln(err)
	}

	// Output:
	// --- Results ---
	// ---
	// | File    |  Line    | Offender |
	// | --------| ---------| -------- |
	// | example.txt | 16 | person@email.com |
	// | example2.txt | 29 | fred@email.com |
}
