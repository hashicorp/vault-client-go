package vault

import (
	"context"
	"fmt"

	"github.com/testcontainers/testcontainers-go"
)

type VaultContainer struct {
	container testcontainers.Container
	address   string
}

func setupVaultContainer(ctx context.Context, rootToken string) (*VaultContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "vault",
		ExposedPorts: []string{"8200/tcp"},
		Env: map[string]string{
			"VAULT_DEV_ROOT_TOKEN_ID": rootToken,
		},
		CapAdd: []string{"IPC_LOCK"},
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to start container: %w", err)
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get container ip: %w", err)
	}

	mappedPort, err := container.MappedPort(ctx, "8200")
	if err != nil {
		return nil, fmt.Errorf("could not get mapped port from container: %w", err)
	}

	vaultContainer := VaultContainer{
		container: container,
		address:   fmt.Sprintf("http://%s:%s", ip, mappedPort.Port()),
	}

	return &vaultContainer, nil
}
