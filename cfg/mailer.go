package cfg

import "github.com/Deirror/dutils/env"

var mailerSuffixes = []string{
	"MAILER_HOST",
	"MAILER_PORT",
	"MAILER_USERNAME",
	"MAILER_PASSWORD",
	"MAILER_FROM",
}

// MailerConfig holds SMTP configuration for sending emails.
type MailerConfig struct {
	Host     string // SMTP server host
	Port     string // SMTP server port
	Username string // SMTP username for authentication
	Password string // SMTP password or token for authentication
	From     string // Default "From" email address
}

func NewMailerConfig(host, port, username, password, from string) *MailerConfig {
	return &MailerConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

// LoadEnvMailerConfig loads MailerConfig values from environment variables,
// supporting an optional prefix.
func LoadEnvMailerConfig(prefix ...string) (*MailerConfig, error) {
	pfx := modPrefix(prefix...)

	host, err := env.GetEnv(pfx + mailerSuffixes[0])
	if err != nil {
		return nil, err
	}

	port, err := env.GetEnv(pfx + mailerSuffixes[1])
	if err != nil {
		return nil, err
	}

	username, err := env.GetEnv(pfx + mailerSuffixes[2])
	if err != nil {
		return nil, err
	}

	password, err := env.GetEnv(pfx + mailerSuffixes[3])
	if err != nil {
		return nil, err
	}

	from, err := env.GetEnv(pfx + mailerSuffixes[4])
	if err != nil {
		return nil, err
	}

	return NewMailerConfig(host, port, username, password, from), nil
}

// LoadEnvMailerConfigs loads multiple MailerConfigs by scanning env vars with mailer suffixes.
func LoadEnvMailerConfigs() (MultiEnvConfig[MailerConfig], error) {
	return LoadMultiEnvConfigs(mailerSuffixes, LoadEnvMailerConfig)
}

// WithHost sets the SMTP host and returns the updated MailerConfig.
func (cfg *MailerConfig) WithHost(host string) *MailerConfig {
	cfg.Host = host
	return cfg
}

// WithPort sets the SMTP port and returns the updated MailerConfig.
func (cfg *MailerConfig) WithPort(port string) *MailerConfig {
	cfg.Port = port
	return cfg
}

// WithUsername sets the SMTP username and returns the updated MailerConfig.
func (cfg *MailerConfig) WithUsername(username string) *MailerConfig {
	cfg.Username = username
	return cfg
}

// WithPassword sets the SMTP password and returns the updated MailerConfig.
func (cfg *MailerConfig) WithPassword(password string) *MailerConfig {
	cfg.Password = password
	return cfg
}

// WithFrom sets the default "From" email address and returns the updated MailerConfig.
func (cfg *MailerConfig) WithFrom(from string) *MailerConfig {
	cfg.From = from
	return cfg
}
