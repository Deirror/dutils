package cfg

import (
	"fmt"
	"strings"

	"github.com/Deirror/dutils/env"
)

// EnvGroupMap maps a prefix (e.g. "WEB", "EXE") to a set of environment variable suffixes and their values.
// For example, for "WEB_JWT_SECRET", the prefix is "WEB" and the suffix is "_JWT_SECRET".
type EnvGroupMap map[string]map[string]string

// Extracts all env vars ending with the suffixes.
func LoadEnvGroups(suffixes []string, filenames ...string) (EnvGroupMap, error) {
	grouped := make(EnvGroupMap)

	envVars, err := env.GetAllEnvs(filenames...)
	if err != nil {
		return nil, err
	}

	for key, val := range envVars {
		for _, suffix := range suffixes {
			if !strings.HasSuffix(key, "_"+suffix) {
				continue
			}

			prefix := strings.TrimSuffix(key, "_"+suffix)

			if len(prefix) == 0 {
				return nil, fmt.Errorf("prefix is empty: _ for suffix: %s", suffix)
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
