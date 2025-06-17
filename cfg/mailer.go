package cfg

import "github.com/Deirror/dutils/env"

type MailerConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
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

func LoadEnvMailerConfig() (*MailerConfig, error) {
	host, err := env.GetEnv("MAILER_HOST")
	if err != nil {
		return nil, err
	}

	port, err := env.GetEnv("MAILER_PORT")
	if err != nil {
		return nil, err
	}

	username, err := env.GetEnv("MAILER_USERNAME")
	if err != nil {
		return nil, err
	}
	password, err := env.GetEnv("MAILER_PASSWORD")
	if err != nil {
		return nil, err
	}

	from, err := env.GetEnv("MAILER_FROM")
	if err != nil {
		return nil, err
	}

	return NewMailerConfig(host, port, username, password, from), nil
}

func (cfg *MailerConfig) WithHost(host string) *MailerConfig {
	cfg.Host = host
	return cfg
}

func (cfg *MailerConfig) WithPort(port string) *MailerConfig {
	cfg.Port = port
	return cfg
}

func (cfg *MailerConfig) WithUsername(username string) *MailerConfig {
	cfg.Username = username
	return cfg
}

func (cfg *MailerConfig) WithPassword(password string) *MailerConfig {
	cfg.Password = password
	return cfg
}

func (cfg *MailerConfig) WithFrom(from string) *MailerConfig {
	cfg.From = from
	return cfg
}
