# dutils

Description
-

Deirrorz Utilities(**dutils**) is a lightweight Go utility library with commonly used helpers and wrappers for building web services. It includes tools for handling environment variables, JSON responses, JWT tokens, password hashing, and more.

## Features

- `env`: Load and validate `.env` files with `godotenv`  
- `json`: Write structured JSON responses with standard error formatting  
- `jwt`: Sign, validate, and manage JWT tokens (including cookie helpers)  
- `crypto`: Securely hash and compare passwords using bcrypt  
- More packages coming soon (e.g., logging, request validation, middleware)


## Packages

### `env`

Helpers for managing environment variables

```go
import "github.com/Deirror/dutils/env"

err := env.InitEnv(".env") // Load .env only if not in production

mode, err := env.GetMode() // "dev" or "prod"

val, err := env.GetEnv("DATABASE_URL")
```

### `crypto`

Password hashing with bcrypt

```go
import "github.com/Deirror/dutils/crypto"

hash, err := crypto.HashPassword("my-secret")
err := crypto.ComparePassword(hash, "my-secret")
```

