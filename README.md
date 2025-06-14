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

### `jwt`

JWT creation, validation, and cookie management

```go
import "github.com/Deirror/dutils/jwt"

jwt.SetJWTCookie(w, token)
jwt.ClearJWTCookie(w)

token, err := jwt.ValidateJWT(tokenString)
```

### `json`

(Coming soon) Standardized JSON error and response writers

## Installation
```bash
go get github.com/Deirror/dutils@latest
```

Or use a specific version:

```bash
go get github.com/Deirror/dutils@v0.1.0
```

## Usage

You can import only the packages you need:

```go
import (
    "github.com/Deirror/dutils/env"
    "github.com/Deirror/dutils/jwt"
)
```

## Roadmap

- Logging wrapper (Zap/Slog)
- JSON response helpers
- Middleware toolkit
- Request validation

## Contributing

Contributions are welcome 🤝! Please open an issue or pull request
