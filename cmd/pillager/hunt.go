// Package pillager contains the command line logic.
//
// The pillager package is the primary consumer of all packages in
// the /pkg directory.
package pillager

import (
	"runtime"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/samsarahq/go/oops"
	"github.com/spf13/cobra"
)

var (
	verbose     bool
	rulesConfig string
	output      string
	templ       string
	workers     int
)

// huntCmd represents the hunt command.
var huntCmd = &cobra.Command{
	Use:   "hunt [directory]",
	Short: "Hunt for loot",
	Long:  "Hunt inside the file system for valuable information",
	Example: `
	Basic:
		pillager hunt .
	
	JSON Format:
		pillager hunt ./example -f json
	
	YAML Format:
		pillager hunt . -f yaml
	
	HTML Format:
		pillager hunt . -f html > results.html
	
	HTML Table Format:
		pillager hunt . -f html-table > results.html
	
	Markdown Table Format:
		pillager hunt . -f table > results.md
	
	Custom Go Template Format:
		pillager hunt . --template "{{ range .Leaks}}Leak: {{.Line}}{{end}}"
	
	Custom Go Template Format from Template File:
		pillager hunt ./example --template "$(cat templates/simple.tmpl)"
`,
	Args: cobra.ExactArgs(1),
	RunE: startHunt(),
}

func init() {
	rootCmd.AddCommand(huntCmd)
	huntCmd.Flags().IntVarP(&workers, "workers", "w", runtime.NumCPU(), "number of concurrent workers to create")
	huntCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "toggle verbose output")
	huntCmd.Flags().StringVarP(&rulesConfig, "rules", "r", "", "path to gitleaks rules.toml config")
	huntCmd.Flags().StringVarP(&output, "format", "f", "json", "set output format (json, yaml)")
	huntCmd.Flags().StringVarP(&templ, "template", "t", "", "set go text/template string for output format")
}

// startHunt builds a Config from the from the given command and arguments,
// then creates a Hunter from the Config and executes it's Hunt function.
func startHunt() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cParams := &config.ConfigParams{
			BasePath:  args[0],
			RulesPath: rulesConfig,
			Format:    config.StringToFormat(output),
			Verbose:   verbose,
			Template:  templ,
			Workers:   workers,
		}

		c, err := config.NewConfig(*cParams)
		if err != nil {
			return oops.Wrapf(err, "creating config")
		}

		h, _ := hunter.NewHunter(*c)
		return h.Hunt()
	}
}
