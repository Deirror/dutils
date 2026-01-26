# servette

>[!WARNING]
> The whole package is in major refactor!
> The documention is out-of-dated

Description
-

***Deirrorz Utilities***(**dutils**) is a lightweight Go utility library with commonly used helpers and wrappers for building web services. It includes tools for handling environment variables, JSON responses, JWT tokens, password hashing, and more.

## Features

- `cfg`: Strongly typed config structs with fluent .With... chaining + multi struct with dynamic prefixes and concrete suffixes
- `crypto`: Secure password hashing and comparison with bcrypt
- `db`: Wrapper for sql.DB with extended funcs
- `env`: Load .env files and manage environment-specific behavior
- `http`: Helpers for Go 1.22+ r.PathValue, query parsing, setting up a std http server, handler wrapper with ctx and custom err type and many more
- `json`: Standardized JSON response and error writers
- `jwt`: Full JWT token generation, validation, and cookie integration
- `logger`: Simple slog-based structured logging setup
- `mail`: Ready to use html templates for sending emails for business purposes
- `oauth`: Google, Facebook and Github OAuth 2.0 config wrappers and more

## Packages

### `cfg`

Typed config structs that can be loaded from environment variables and modified via fluent .With... chaining.

Supported configs:
- cfg.BlobConfig
- cfg.PaymentConfig
- cfg.DBConfig
- cfg.JWTConfig
- cfg.KVConfig
- cfg.MailerConfig
- cfg.OAuthConfig
- and others

```go
import "github.com/Deirror/dutils/cfg"

cfg, err := cfg.LoadEnvMailerConfig()
// .With... pattern to override specific fields
cfg.WithHost("smtp.newhost.com").WithPort("465")
```

### `env`

Helpers for loading environment variables safely.

```go
import "github.com/Deirror/dutils/env"

_ = env.LoadEnv(".env") // Load only if not in production
mode, _ := env.GetMode() // "dev" or "prod"

val, err := env.GetEnv("DATABASE_URL")
```

### `crypto`

Secure password hashing and verification.

```go
import "github.com/Deirror/dutils/crypto"

hash, err := crypto.HashPassword("password123")
err = crypto.ComparePassword(hash, "password123")
```

### `http`

Simplified helpers for parsing Go 1.22+ path and query values.

```go
import "github.com/Deirror/dutils/http"

s := http.NewStdServerFromConfig(cfg, h)      
```

### `jwt`

JWT handling with token signing, validation, and cookie management.

```go
import "github.com/Deirror/dutils/jwt"

jwtHandler := jwt.NewJWT(jwtCfg)

token, err := jwtHandler.GenerateToken("user@example.com")

jwtHandler.SetCookie(w, token)

email, err := jwtHandler.ValidateJWT(tokenFromRequest)

jwt.SetJWTCookie(w, token)  // raw cookie helper
jwt.ClearJWTCookie(w)       // clear cookie
```

### `logger`

Minimal structured logger using Go's standard slog.

```go
import "github.com/Deirror/dutils/logger"

log := logger.Init("dev") // or "prod"
log.Info("app started", "version", "v1.0.0")
```

### `json`

Helpers for reading and writing structured JSON in handlers and services.

```go
import "github.com/Deirror/dutils/json"
```

- Response Writers
  
```go
err := json.Write(w, http.StatusOK, myData) // Marshal and write JSON
```

- Request Parsers

```go
var data MyStruct
err := json.ParseJSONInto(r.Body, &data)

data, err := json.DecodeInto[MyStruct](r.Body) // Generics-based shortcut
```

- Encoder (manual streams)

```go
err := json.Encode(w, myStruct) // Stream JSON to any io.Writer
```

### `db`

A wrapper struct for sql.DB, which has Close and Ping funcs.

```go
db, err := Connect(cfg.Driver, cfg.DSN) // Uses db config
if err != nil {
    return nil, err
}

// Sets parameters to connection pool
db.SetMaxOpenConns(int(cfg.PoolSize))
db.SetMaxIdleConns(int(cfg.MaxIdle))
db.SetConnMaxLifetime(cfg.MaxLifetime)
```

## Installation
```bash
go get github.com/Deirror/dutils@latest
```

Or use a specific version:

```bash
go get github.com/Deirror/dutils@latest
```

## Usage

You can import only the packages you need:

```go
import (
    "github.com/Deirror/dutils/env"
    "github.com/Deirror/dutils/jwt"
)
```

## Contributing

Contributions are welcome 🤝! Please open an issue or pull request
