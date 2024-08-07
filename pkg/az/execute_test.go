package az

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path"
	"testing"

	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)

	// Validate that we passed in the user agent environment variable when we called the az cli
	_, hasEnv := os.LookupEnv(AzureUserAgentEnvVar)
	if !hasEnv {
		fmt.Println("expected the az cli to be called with the AZURE_HTTP_USER_AGENT environment variable set")
		os.Exit(127)
	}
}

func TestMixin_Execute(t *testing.T) {
	testcases := []struct {
		name        string
		file        string
		wantOutput  string
		wantCommand string
	}{
		{"install", "testdata/install-input.yaml", "",
			"az login --output json --password password --service-principal --tenant tenant --username client-id"},
		{"install group", "testdata/install-group.yaml", "",
			"az group create --location westus --name mygroup"},
		{"update group", "testdata/upgrade-group.yaml", "",
			"az group create --location westus --name mygroup"},
		{"uninstall group", "testdata/uninstall-group.yaml", "",
			"az group delete --yes --name mygroup"},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			m := NewTestMixin(t)

			m.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := os.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute(ctx)
			require.NoError(t, err, "execute failed")

			if tc.wantOutput == "" {
				outputs, _ := m.FileSystem.ReadDir("/cnab/app/porter/outputs")
				assert.Empty(t, outputs, "expected no outputs to be created")
			} else {
				wantPath := path.Join("/cnab/app/porter/outputs", tc.wantOutput)
				exists, _ := m.FileSystem.Exists(wantPath)
				assert.True(t, exists, "output file was not created %s", wantPath)
			}
		})
	}
}
