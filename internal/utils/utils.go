package utils

func EmptyString() *string {
	var s string
	return &s
}

func ReplaceNilWithEmptyString(s **string) {
	*s = EmptyString()
}
