package hunter_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/brittonhayes/pillager/pkg/hunter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseRulesFromConfigFile(t *testing.T) {
	testPillagerConfigStr := `
	title = "pillager config"
	[[rules]]
		description = "email"
		regex = '''^[A-Z0-9_!#$%&'*+/=?{|}~^.-]+@[A-Z0-9.-]+$'''
		tags = ["email"]
	`
	tmpValidConfig, err := ioutil.TempFile("", "toml-")
	require.NoError(t, err)
	defer os.Remove(tmpValidConfig.Name())

	_, err = tmpValidConfig.WriteString(testPillagerConfigStr)
	require.NoError(t, err)
	err = tmpValidConfig.Close()
	require.NoError(t, err)

	testcases := []struct {
		desc             string
		filepath         string
		errExpected      bool
		numRulesExpected int
	}{
		{
			desc:        "invalid custom filepath",
			filepath:    "bad/file/path",
			errExpected: true,
		},
		{
			desc:             "valid custom filepath",
			filepath:         tmpValidConfig.Name(),
			errExpected:      false,
			numRulesExpected: 1,
		},
		{
			desc:             "no filepath provided",
			filepath:         "",
			errExpected:      false,
			numRulesExpected: 6,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			gitleaks, err := hunter.ParseRulesFromConfigFile(tc.filepath)

			if tc.errExpected {
				assert.Error(t, err)
				assert.Nil(t, gitleaks)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, gitleaks)
				assert.Len(t, gitleaks.Rules, tc.numRulesExpected)
			}
		})
	}
}
