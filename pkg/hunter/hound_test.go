package hunter

import (
	"github.com/brittonhayes/pillager/templates"
	"github.com/zricethezav/gitleaks/v7/scan"
)

// Here is an example of utilizing the Howl function on a slice
// of findings. The Howl method is the final method in the
// hunting process. It takes whatever has been found and
// outputs it for the user.
func ExampleHound_Howl_json() {
	h := NewHound(CustomFormat, &templates.Table)

	h.Findings = &scan.Report{
		Leaks: []scan.Leak{
			{
				Line:       "person@email.com",
				LineNumber: 16,
				Offender:   "person@email.com",
				Rule:       "Email Addresses",
				File:       "example.txt",
			},
			{
				Line:       "fred@email.com",
				LineNumber: 29,
				Offender:   "fred@email.com",
				Rule:       "Email Addresses",
				File:       "example2.txt",
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
