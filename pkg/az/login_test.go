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
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_CLIENT_SECRET", "test-client-secret")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_TENANT_ID", "test-tenant-id")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	defer func() {
		err := os.Unsetenv("AZURE_CLIENT_ID")
		if err != nil {
			t.Fatal("failed to unset env", err)
		}
		err = os.Unsetenv("AZURE_CLIENT_SECRET")
		if err != nil {
			t.Fatal("failed to unset env", err)
		}
		err = os.Unsetenv("AZURE_TENANT_ID")
		if err != nil {
			t.Fatal("failed to unset env", err)
		}
	}()

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{"login"}
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetCommandAndGetArguments_ExistingAzureDirectory(t *testing.T) {
	tempHome := t.TempDir()
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	homeDir := os.Getenv("HOME")
	if err := os.MkdirAll(filepath.Join(homeDir, ".azure"), 0o755); err != nil {
		t.Fatal("failed to create .azure directory:", err)
	}
	defer func() {
		err := os.RemoveAll(filepath.Join(homeDir, ".azure"))
		if err != nil {
			t.Fatal("failed to tidy up", err)
		}
	}()

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{}
	assert.Equal(t, "true", cmd.GetCommand())
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetArguments_ManagedIdentity(t *testing.T) {
	tempHome := t.TempDir()
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Unsetenv("AZURE_CLIENT_ID")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Unsetenv("AZURE_CLIENT_SECRET")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Unsetenv("AZURE_TENANT_ID")
	if err != nil {
		t.Fatal("failed to set env", err)
	}

	cmd := &LoginCommand{}
	args := cmd.GetArguments()

	expectedArgs := []string{"login"}
	assert.Equal(t, expectedArgs, args)
}

func TestLoginCommand_GetFlags_ServicePrincipal(t *testing.T) {
	tempHome := t.TempDir()
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_CLIENT_SECRET", "test-client-secret")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_TENANT_ID", "test-tenant-id")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	defer func() {
		err := os.Unsetenv("AZURE_CLIENT_ID")
		if err != nil {
			t.Fatal("failed to set env", err)
		}
		err = os.Unsetenv("AZURE_CLIENT_SECRET")
		if err != nil {
			t.Fatal("failed to set env", err)
		}
		err = os.Unsetenv("AZURE_TENANT_ID")
		if err != nil {
			t.Fatal("failed to set env", err)
		}
	}()

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
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	err = os.Setenv("AZURE_CLIENT_ID", "test-client-id")
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	defer func() {
		err := os.Unsetenv("AZURE_CLIENT_ID")
		if err != nil {
			t.Fatal("unable to unset env", err)
		}
	}()

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
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{
		builder.NewFlag("identity", ""),
	}
	assert.Equal(t, expectedFlags, flags)
}

func TestLoginCommand_GetFlags_ExistingAzureDirectory(t *testing.T) {
	tempHome := t.TempDir()
	err := os.Setenv("HOME", tempHome)
	if err != nil {
		t.Fatal("failed to set env", err)
	}
	homeDir := os.Getenv("HOME")
	if err := os.MkdirAll(filepath.Join(homeDir, ".azure"), 0o755); err != nil {
		t.Fatal("failed to create .azure directory:", err)
	}
	defer func() {
		err := os.RemoveAll(filepath.Join(homeDir, ".azure"))
		if err != nil {
			t.Fatal("failed to tidy up", err)
		}
	}()

	cmd := &LoginCommand{}
	flags := cmd.GetFlags()

	expectedFlags := builder.Flags{}
	assert.Equal(t, expectedFlags, flags)
}
