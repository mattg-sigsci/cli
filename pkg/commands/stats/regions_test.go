package stats_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/fastly/cli/pkg/app"
	"github.com/fastly/cli/pkg/mock"
	"github.com/fastly/cli/pkg/testutil"
	"github.com/fastly/go-fastly/v5/fastly"
)

func TestRegions(t *testing.T) {
	args := testutil.Args
	for _, testcase := range []struct {
		args       []string
		api        mock.API
		wantError  string
		wantOutput string
	}{
		{
			args:       args("stats regions"),
			api:        mock.API{GetRegionsFn: getRegionsOK},
			wantOutput: "foo\nbar\nbaz\n",
		},
		{
			args:      args("stats regions"),
			api:       mock.API{GetRegionsFn: getRegionsError},
			wantError: errTest.Error(),
		},
	} {
		t.Run(strings.Join(testcase.args, " "), func(t *testing.T) {
			var stdout bytes.Buffer
			opts := testutil.NewRunOpts(testcase.args, &stdout)
			opts.APIClient = mock.APIClient(testcase.api)
			err := app.Run(opts)
			testutil.AssertErrorContains(t, err, testcase.wantError)
			testutil.AssertStringContains(t, stdout.String(), testcase.wantOutput)
		})
	}
}

func getRegionsOK() (*fastly.RegionsResponse, error) {
	return &fastly.RegionsResponse{
		Data: []string{"foo", "bar", "baz"},
	}, nil
}

var errTest = errors.New("fixture error")

func getRegionsError() (*fastly.RegionsResponse, error) {
	return nil, errTest
}
