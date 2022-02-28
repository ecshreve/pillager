package pillager

import (
	"log"

	"github.com/gookit/color"
	"github.com/mitchellh/go-homedir"
	"github.com/samsarahq/go/oops"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "pillager",
	Short: "Pillage systems for sensitive information",
	Long: color.Cyan.Text(`
	██▓███   ██▓ ██▓     ██▓    ▄▄▄        ▄████ ▓█████  ██▀███
	▓██░  ██▒▓██▒▓██▒    ▓██▒   ▒████▄     ██▒ ▀█▒▓█   ▀ ▓██ ▒ ██▒
	▓██░ ██▓▒▒██▒▒██░    ▒██░   ▒██  ▀█▄  ▒██░▄▄▄░▒███   ▓██ ░▄█ ▒
	▒██▄█▓▒ ▒░██░▒██░    ▒██░   ░██▄▄▄▄██ ░▓█  ██▓▒▓█  ▄ ▒██▀▀█▄
	▒██▒ ░  ░░██░░██████▒░██████▒▓█   ▓██▒░▒▓███▀▒░▒████▒░██▓ ▒██▒
	▒▓▒░ ░  ░░▓  ░ ▒░▓  ░░ ▒░▓  ░▒▒   ▓▒█░ ░▒   ▒ ░░ ▒░ ░░ ▒▓ ░▒▓░
	░▒ ░      ▒ ░░ ░ ▒  ░░ ░ ▒  ░ ▒   ▒▒ ░  ░   ░  ░ ░  ░  ░▒ ░ ▒░
	░░        ▒ ░  ░ ░     ░ ░    ░   ▒   ░ ░   ░    ░     ░░   ░
	░      ░  ░    ░  ░     ░  ░      ░    ░  ░   ░

			Pillage filesystems for loot.
`),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by pillager.pillager(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(oops.Wrapf(err, "encountered problem executing root command"))
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pillager.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			log.Fatalln(oops.Wrapf(err, "encountered problem finding home directory"))
		}

		// Search for config in home directory with name ".pillager" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".pillager")
	}

	viper.AutomaticEnv() // Read in environment variables that match.

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("using config file:", viper.ConfigFileUsed())
	}
}
