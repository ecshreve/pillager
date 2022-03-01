package config_test

import (
	"testing"

	"github.com/brittonhayes/pillager/pkg/config"
	"github.com/stretchr/testify/assert"
	gitleaks "github.com/zricethezav/gitleaks/v8/config"
)

func TestValidate(t *testing.T) {
	testcases := []struct {
		desc        string
		cfg         *config.Cfg
		errExpected bool
	}{
		{
			desc:        "valid path",
			cfg:         &config.Cfg{".", false, 1, 1, "", "", &gitleaks.Config{Rules: []*gitleaks.Rule{{Description: "test"}}}},
			errExpected: false,
		},
		{
			desc:        "invalid path",
			cfg:         &config.Cfg{"adsfasd", false, 1, 1, "", "", &gitleaks.Config{Rules: []*gitleaks.Rule{{Description: "test"}}}},
			errExpected: true,
		},
		{
			desc:        "invalid format",
			cfg:         &config.Cfg{".", false, 1, 7777, "", "", &gitleaks.Config{Rules: []*gitleaks.Rule{{Description: "test"}}}},
			errExpected: true,
		},
		{
			desc:        "invalid workers",
			cfg:         &config.Cfg{".", false, 777, 1, "", "", &gitleaks.Config{Rules: []*gitleaks.Rule{{Description: "test"}}}},
			errExpected: true,
		},
		{
			desc:        "invalid rules",
			cfg:         &config.Cfg{".", false, 1, 1, "", "", nil},
			errExpected: true,
		},
		{
			desc:        "invalid nested rule",
			cfg:         &config.Cfg{".", false, 1, 1, "", "", &gitleaks.Config{}},
			errExpected: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.desc, func(t *testing.T) {
			err := testcase.cfg.Validate()

			if testcase.errExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseRules(t *testing.T) {
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
			filepath:         "./testdata/test.rules.toml",
			errExpected:      false,
			numRulesExpected: 95,
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
			rules, err := config.ParseRules(tc.filepath)

			if tc.errExpected {
				assert.Error(t, err)
				assert.Nil(t, rules)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, rules)
				assert.Len(t, rules.Rules, tc.numRulesExpected)
			}
		})
	}
}
