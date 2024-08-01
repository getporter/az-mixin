package az

import (
	"bytes"
	"context"
	"os"
	"testing"

	"get.porter.sh/mixin/az/pkg"
	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	testcases := []struct {
		name           string
		inputFile      string
		wantOutputFile string
	}{
		{name: "build with config", inputFile: "testdata/build-input-with-config.yaml", wantOutputFile: "testdata/build-with-config.txt"},
		{name: "build without config", inputFile: "testdata/build-input-without-config.yaml", wantOutputFile: "testdata/build-without-config.txt"},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			// Set a fake version of the mixin and porter for our user agent
			pkg.Version = "v1.2.3"

			b, err := os.ReadFile(tc.inputFile)
			require.NoError(t, err)

			m := NewTestMixin(t)
			m.In = bytes.NewReader(b)

			err = m.Build(context.Background())
			require.NoError(t, err, "build failed")

			test.CompareGoldenFile(t, tc.wantOutputFile, m.TestContext.GetOutput())
		})
	}
}
