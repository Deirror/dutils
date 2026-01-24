package envcfg

// Represents maps of  prefixes (e.g. "NEON", "UPSTASH", "WEB") to a set of Configs.
type MultiConfig[T any] map[string]*T

// Template func for loading envs with their prefixes.
func LoadMultiConfig[T any](
	suffixes []string,
	loader func(prefix ...string) (*T, error),
) (MultiConfig[T], error) {
	grouped, err := LoadGroups(suffixes)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*T)
	for prefix := range grouped {
		conf, err := loader(prefix)
		if err != nil {
			return nil, err
		}
		result[prefix] = conf
	}

	return result, nil
}
