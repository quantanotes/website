package utils

func EmptyString() *string {
	var s string
	return &s
}

func EmptyMap() map[string]any {
	return make(map[string]any)
}

func ReplaceNilWithEmptyString(s **string) {
	if *s == nil {
		*s = EmptyString()
	}
}

func ReplaceNilWithEmptyMap(m *map[string]any) {
	if *m == nil {
		*m = EmptyMap()
	}
}
