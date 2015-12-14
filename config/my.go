package config

var store map[string]string

func init() {
	store = map[string]string{}
}


func Set(key string, val string) {
	store[key] = val
}

func Get(key string) string{
	if v, exists := store[key]; exists {
		return v
	}
	return ""
}