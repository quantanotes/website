package utils

func EmptyString() *string {
	var s string
	return &s
}

func ReplaceNilWithEmptyString(s **string) {
	if *s == nil {
		*s = EmptyString()
	}
}
