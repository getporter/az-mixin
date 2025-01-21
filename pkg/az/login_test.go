package az

import (
	"os"
	"path/filepath"
	"testing"

	"get.porter.sh/porter/pkg/exec/builder"
	"github.com/stretchr/testify/assert"
)

func TestLoginCommand_GetArguments_ServicePrincipal(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	os.Setenv("AZURE_CLIENT_SECRET", "test-client-secret")
	os.Setenv("AZURE_TENANT_ID", "test-tenant-id")
	defer os.Unsetenv("AZURE_CLIENT_ID")
	defer os.Unsetenv("AZURE_CLIENT_SECRET")
	defer os.Unsetenv("AZURE_TENANT_ID")

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{"login"}
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetCommandAndGetArguments_ExistingAzureDirectory(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	homeDir := os.Getenv("HOME")
	os.MkdirAll(filepath.Join(homeDir, ".azure"), 0755)
	defer os.RemoveAll(filepath.Join(homeDir, ".azure"))

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{}
	assert.Equal(t, "true", cmd.GetCommand())
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetArguments_ManagedIdentity(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	os.Unsetenv("AZURE_CLIENT_ID")
	os.Unsetenv("AZURE_CLIENT_SECRET")
	os.Unsetenv("AZURE_TENANT_ID")

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{"login"}
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetFlags_ServicePrincipal(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	os.Setenv("AZURE_CLIENT_SECRET", "test-client-secret")
	os.Setenv("AZURE_TENANT_ID", "test-tenant-id")
	defer os.Unsetenv("AZURE_CLIENT_ID")
	defer os.Unsetenv("AZURE_CLIENT_SECRET")
	defer os.Unsetenv("AZURE_TENANT_ID")

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{
		builder.NewFlag("service-principal", ""),
		builder.NewFlag("username", "test-client-id"),
		builder.NewFlag("password", "test-client-secret"),
		builder.NewFlag("tenant", "test-tenant-id"),
	}
	assert.Equal(t, expectedFlags, flags)
}

func TestLoginCommand_GetFlags_UserAssignedManagedIdentity(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	defer os.Unsetenv("AZURE_CLIENT_ID")

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{
		builder.NewFlag("identity", ""),
		builder.NewFlag("username", "test-client-id"),
	}
	assert.Equal(t, expectedFlags, flags)
}

func TestLoginCommand_GetFlags_SystemManagedIdentity(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{
		builder.NewFlag("identity", ""),
	}
	assert.Equal(t, expectedFlags, flags)
}

func TestLoginCommand_GetFlags_ExistingAzureDirectory(t *testing.T) {
	tempHome := t.TempDir()
	os.Setenv("HOME", tempHome)
	homeDir := os.Getenv("HOME")
	os.MkdirAll(filepath.Join(homeDir, ".azure"), 0755)
	defer os.RemoveAll(filepath.Join(homeDir, ".azure"))

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{}
	assert.Equal(t, expectedFlags, flags)
}
