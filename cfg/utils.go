package cfg

import "strings"

func modPrefix(prefix ...string) string {
	pfx := ""
	if len(prefix) > 0 && prefix[0] != "" {
		pfx = prefix[0]
		if !strings.HasSuffix(pfx, "_") {
			pfx += "_"
		}
	}
	return pfx
}
