package commands

import (
	"bytes"
	"fmt"
	"testing"
)

func TestHelp(t *testing.T) {
	cmd := rootCmd
	buf := new(bytes.Buffer)
	cmd.SetOutput(buf)
	cmd.SetArgs([]string{
		"hunt",
		"./testdata",
		"-f",
		"yaml",
		"--rules",
		"./testdata/rulestest.toml",
	})
	err := cmd.Execute()
	if err != nil {
		fmt.Println("bad")
	}
	fmt.Println(buf)
}
