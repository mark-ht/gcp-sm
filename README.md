# GCP-SM
One liner helper for gcp sm to retrieve latest secret value.

## Usage

```go
ctx := context.Background()
secret, err := GetSecret(ctx, "projects/my-project/secrets/my-secret")
if err != nil {
    log.Fatalf("Error retrieving secret: %v", err)
}
fmt.Println("Retrieved secret:", secret)
```

## Installation

```sh
go get github.com/mark-ht/gcp-sm
```

## Authentication
Ensure you have necessary creds set on env to access target secret.