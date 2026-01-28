# servette

Description
-

***Servette*** is a lightweight Go utility library with commonly used helpers and wrappers for building web services. It includes tools for app lifecycle, transport helpers, handling environment variables, JSON responses, JWT tokens, password hashing, and more

Overall philosophy - remove boiler plate even further and just use already existing code and knowledge to build faster programs without thinking too much about the process

Main architecture follows this pattern:
- `config` - consists of parsing data such as client creditianls for any http service, app related and more
  - `env` - most used way to pass values to a Go struct via `.env` file outside the code itself. Parsing has some hardcoded suffixes and dynamic prefixes, making it kind of flexible. For example, suffix for **JWT** is `JWT_COOKIE`, but the preifx is you choice - `WEB_`, etc. 
- `logger` - log is used all over the place. It has some timing functions and can be used for the logger service pattern
- `transport` - the way you want to transfer the data and run handlers and so on. ALso forces to use a custom error type which is more web-oriented/more details
  - `http` - most used way to comunicate. A handler func type is defined with a `ctx` var in the func signature 
- `app` - overall program lifecycle. Runs the servers by having an interface `Runner` which allow to run any kind of server or whatever you consider as a runner

## Contributing

Contributions are welcome 🤝! Please open an issue or pull request
