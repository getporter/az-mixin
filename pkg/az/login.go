package az

import (
	"context"
	"os"
	"path/filepath"

	"get.porter.sh/porter/pkg/exec/builder"
)

var (
	_ TypedCommand             = &LoginCommand{}
	_ builder.HasErrorHandling = &LoginCommand{}
)

// LoginCommand handles logging into Azure
type LoginCommand struct {
	action      string
	Description string `yaml:"description"`
}

func (c *LoginCommand) HandleError(ctx context.Context, err builder.ExitError, stdout string, stderr string) error {
	// Handle specific login errors if necessary
	return err
}

func (c *LoginCommand) GetWorkingDir() string {
	return ""
}

func (c *LoginCommand) SetAction(action string) {
	c.action = action
}

func (c *LoginCommand) GetCommand() string {
	if c.azureDirExists() {
		// Use a no-op command since we don't have to log in.
		return "true"
	}

	return "az"
}

func (c *LoginCommand) GetArguments() []string {
	if c.azureDirExists() {
		return []string{}
	}
	return []string{"login"}
}

func (c *LoginCommand) GetFlags() builder.Flags {
	flags := builder.Flags{}

	if c.azureDirExists() {
		return flags
	}

	if os.Getenv("AZURE_CLIENT_ID") != "" && os.Getenv("AZURE_CLIENT_SECRET") != "" && os.Getenv("AZURE_TENANT_ID") != "" {
		// Add flags for service principal authentication
		flags = append(flags, builder.NewFlag("service-principal", ""))
		flags = append(flags, builder.NewFlag("username", os.Getenv("AZURE_CLIENT_ID")))
		flags = append(flags, builder.NewFlag("password", os.Getenv("AZURE_CLIENT_SECRET")))
		flags = append(flags, builder.NewFlag("tenant", os.Getenv("AZURE_TENANT_ID")))
	} else if os.Getenv("AZURE_CLIENT_ID") != "" {
		// Add flag for user-assigned managed identity
		flags = append(flags, builder.NewFlag("identity", ""))
		flags = append(flags, builder.NewFlag("username", os.Getenv("AZURE_CLIENT_ID")))
	} else {
		// Add flag for system-assigned managed identity
		flags = append(flags, builder.NewFlag("identity", ""))
	}

	return flags
}

func (c *LoginCommand) SuppressesOutput() bool {
	return false
}

func (c *LoginCommand) azureDirExists() bool {
	_, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".azure"))
	return err == nil
}
