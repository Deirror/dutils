package envcfg

// Represents maps of  prefixes (e.g. "NEON", "UPSTASH", "WEB") to a set of EnvConfigs.
type MultiEnvConfig[T any] map[string]*T

// Template func for loading envs with their prefixes.
func LoadMultiEnvConfigs[T any](
	suffixes []string,
	loader func(prefix ...string) (*T, error),
) (MultiEnvConfig[T], error) {
	grouped, err := LoadEnvGroups(suffixes)
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
