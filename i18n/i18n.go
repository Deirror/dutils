package i18n

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/Deirror/servette/encoding/json"
)

// lang -> key -> val
type Bundle map[string]map[string]string

var i18n Bundle

func Load(dir string) error {
	i18n = make(Bundle)

	files, _ := os.ReadDir(dir)
	for _, f := range files {
		lang := strings.TrimSuffix(f.Name(), ".json")

		data, err := os.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			return err
		}

		r := bytes.NewReader(data)
		msgs, err := json.Decode[map[string]string](r)
		if err != nil {
			return err
		}

		i18n[lang] = msgs
	}
	return nil

}

func Tr(ctx context.Context, key string) string {
	lang, _ := ctx.Value(LangKey).(string)

	if m, ok := i18n[lang][key]; ok {
		return m
	}

	// fallback to English
	if m, ok := i18n[DefaultLang][key]; ok {
		return m
	}

	return key //debugging-friendly fallback
}
