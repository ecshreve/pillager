package hunter

import (
	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/templates"
	"github.com/zricethezav/gitleaks/v8/report"
)

// Here is an example of utilizing the Howl function on a slice
// of findings. The Howl method is the final method in the
// hunting process. It takes whatever has been found and
// outputs it for the user.
func ExampleHound_Howl_json() {
	h := NewHound(config.CustomFormat, &templates.Table)

	h.Findings = &Report{
		Leaks: []report.Finding{
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

	h.Howl()

	// Output:
	// ---
	// Hooooowl -- 🐕
	// ---
	// | File    |  Line    | Offender |
	// | --------| ---------| -------- |
	// | example.txt | 16 | person@email.com |
	// | example2.txt | 29 | fred@email.com |
}
