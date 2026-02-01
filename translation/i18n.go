package translation

type I18n struct {
	data map[string]string
}


func (i *I18n) T(key string) string {
	if v, ok := i.data[key]; ok {
		return v
	}
	return key
}

