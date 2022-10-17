package vault

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

type VaultContainer struct {
	testcontainers.Container
	address string
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
		Container: container,
		address:   fmt.Sprintf("http://%s:%s", ip, mappedPort.Port()),
	}

	return &vaultContainer, nil
}

func TestClient_SimpleReadWrite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration tests")
	}

	const rootToken = "read&write"
	ctx := context.Background()

	vault, err := setupVaultContainer(ctx, rootToken)
	require.NoError(t, err)
	defer vault.Terminate(ctx)

	client, err := NewClient(Configuration{
		BaseAddress: vault.address,
	})
	require.NoError(t, err)

	require.NoError(t, client.SetToken(rootToken))

	const secretPath = "/secret/data/my-secret"
	expected := map[string]interface{}{
		"password1": "abc123",
		"password2": "trustno1",
	}

	_, err = client.Write(ctx, secretPath, map[string]interface{}{
		"data": expected, // TODO: is this key always `data`?
	})
	require.NoError(t, err)

	resp, err := client.Read(ctx, secretPath)
	require.NoError(t, err)

	// TODO: can we assert other things about the response here?
	require.Equal(t, expected, resp.Data["data"])
}
