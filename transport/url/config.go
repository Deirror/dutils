package urlx

// Config holds configuration for an URL.
type Config struct {
	URL string // any service URL
}

func NewConfig(url string) *Config {
	return &Config{
		URL: url,
	}
}

// WithURL sets the URL field and returns the updated config.
func (c *Config) WithURL(url string) *Config {
	c.URL = url
	return c
}
