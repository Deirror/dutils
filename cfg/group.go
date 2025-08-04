package cfg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Deirror/dutils/env"
)

// EnvGroupMap maps a prefix (e.g. "WEB", "EXE") to a set of environment variable suffixes and their values.
// For example, for "WEB_JWT_SECRET", the prefix is "WEB" and the suffix is "_JWT_SECRET".
type EnvGroupMap map[string]map[string]string

// Extracts all env vars ending with the suffixes.
func LoadEnvGroups(suffixes []string) (EnvGroupMap, error) {
	grouped := make(EnvGroupMap)

	envVars, err := env.GetAllEnv()
	if err != nil {
		return nil, err
	}

	for _, envVar := range envVars {
		parts := strings.SplitN(envVar, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("cannot split env var: %v", parts)
		}

		key := parts[0]
		val := parts[1]

		for _, suffix := range suffixes {
			if !strings.HasSuffix(key, suffix) {
				continue
			}

			prefixWithUnderscore := strings.TrimSuffix(key, suffix)

			if !strings.HasSuffix(prefixWithUnderscore, "_") {
				return nil, fmt.Errorf("%s must end in _", prefixWithUnderscore)
			}

			prefix := strings.TrimSuffix(prefixWithUnderscore, "_")

			if len(prefix) == 0 {
				return nil, errors.New("prefix is empty: _")
			}

			if _, ok := grouped[prefix]; !ok {
				grouped[prefix] = make(map[string]string)
			}
			grouped[prefix][suffix] = val
		}
	}

	return grouped, nil
}

// Retrieves env vars based on prefix.
func (e EnvGroupMap) GetGroup(prefix string) map[string]string {
	return e[prefix]
}
