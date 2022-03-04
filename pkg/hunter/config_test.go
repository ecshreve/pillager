package hunter_test

import (
	"fmt"
	"testing"

	"github.com/samsarahq/go/snapshotter"
	"github.com/zricethezav/gitleaks/v8/config"

	"github.com/brittonhayes/pillager/pkg/format"
	"github.com/brittonhayes/pillager/pkg/hunter"
)

func TestConfig(t *testing.T) {
	snapshotter := snapshotter.New(t)
	snapshotter.SnapshotErrors = true
	defer snapshotter.Verify()

	testcases := []struct {
		desc string
		cfg  *hunter.Config
	}{
		{
			desc: "default",
			cfg:  hunter.NewConfig(),
		},
		{
			desc: "verbose true",
			cfg:  hunter.NewConfig(hunter.WithVerbose(true)),
		},
		{
			desc: "set some fields",
			cfg: hunter.NewConfig(
				hunter.WithScanPath("./testdata"),
				hunter.WithWorkers(3),
				hunter.WithVerbose(true),
				hunter.WithRedact(true),
				hunter.WithFormat(format.StringToReporter("json")),
				hunter.WithLogLevel("debug"),
			),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			// Remove the gitleaks field before snapshotting for readability.
			tc.cfg.Gitleaks = config.Config{}
			snapshotter.Snapshot(tc.desc, tc.cfg)
			fmt.Printf("%+v\n", tc.cfg)
		})
	}
}
