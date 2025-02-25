package vault

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

// Vault represents a vault struct used for reading and writing secrets.
type Vault struct {
	Client *api.Client
}

// NewDefaultClient returns a new vault client wrapper.
func NewDefaultClient() (*Vault, error) {
	_, ok := os.LookupEnv("VAULT_ADDR")
	if !ok {
		return nil, fmt.Errorf("VAULT_ADDR required but not set")
	}

	vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
	if !ok {
		return nil, fmt.Errorf("VAULT_TOKEN required but not set")
	}

	config := api.DefaultConfig()
	if err := config.ReadEnvironment(); err != nil {
		return nil, err
	}

	c, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	c.SetToken(vaultToken)

	vaultNamespace, ok := os.LookupEnv("VAULT_NAMESPACE")
	if ok {
		c.SetNamespace(vaultNamespace)
	}

	return &Vault{Client: c}, nil
}

// NewClient returns a new vault client wrapper.
func NewClient(addr, token string) (*Vault, error) {
	cfg := &api.Config{
		Address: addr,
	}

	c, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	c.SetToken(token)

	return &Vault{Client: c}, nil
}
