package gcpsm

import (
	"context"
	"fmt"
	"strings"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	secretmanagerpb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

// GetSecret retrieves the secret value from Google Cloud Secret Manager.
//
// This function takes the secret's resource path as input and returns the latest version's secret value as a string.
//
// Parameters:
//   - ctx: The context for managing request deadlines and cancellation.
//   - secretPath: The resource path of the secret in the format
//     "projects/{project}/secrets/{secret}/versions/{version}".
//
// Returns:
//   - A *secretmanagerpb.SecretPayload containing the secret data.
//   - An error if occurs
func GetSecret(ctx context.Context, secretPath string) (*secretmanagerpb.SecretPayload, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create secretmanager client: %w", err)
	}
	defer client.Close()

	if !strings.HasSuffix(secretPath, "/versions/latest") {
		secretPath = fmt.Sprintf("%s/versions/latest", secretPath)
	}

	req := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretPath,
	}

	result, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to access secret version: %w", err)
	}

	return result.Payload, nil
}
