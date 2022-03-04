package rules_test

import (
	"testing"

	"github.com/brittonhayes/pillager/pkg/rules"
	"github.com/samsarahq/go/snapshotter"
)

func TestRulesLoader(t *testing.T) {
	snapshotter := snapshotter.New(t)
	snapshotter.SnapshotErrors = true
	defer snapshotter.Verify()

	testcases := []struct {
		desc   string
		loader *rules.Loader
	}{
		{
			desc:   "simple",
			loader: rules.NewLoader(),
		},
		{
			desc:   "strict",
			loader: rules.NewLoader(rules.WithStrict()),
		},
		{
			desc:   "from file",
			loader: rules.NewLoader(rules.FromFile("./testdata/rulestest.toml")),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			ll := tc.loader.Load()
			snapshotter.Snapshot(tc.desc, ll)
		})
	}

}
