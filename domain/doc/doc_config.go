package doc

// Config holds configuration details for connecting to a document database.
type Config struct {
	StoreURL string // used such as a MongoDB connection string (example: mongodb+srv://.../dbName)
}

func NewConfig(storeURL string) *Config {
	return &Config{
		StoreURL: storeURL,
	}
}

// WithStoreURL sets a new document store URL and returns the updated Config.
func (c *Config) WithStoreURL(url string) *Config {
	c.StoreURL = url
	return c
}
