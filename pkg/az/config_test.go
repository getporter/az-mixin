package az

import (
	"testing"

	"get.porter.sh/mixin/az/pkg"
	"github.com/stretchr/testify/require"
)

func TestSetUserAgent(t *testing.T) {
	pkg.Commit = "abc123"
	pkg.Version = "v1.2.3"

	m := NewTestMixin(t)
	m.SetUserAgent()

	expected := "porter az/v1.2.3"
	require.Equal(t, expected, m.Getenv(AZURE_HTTP_USER_AGENT))
}
