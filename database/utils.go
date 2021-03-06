package database

func String(s string) *string {
	return &s
}
func StringVal(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func Int(i int) *int {
	return &i
}
func IntVal(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

func Bool(b bool) *bool {
	return &b
}
func BoolVal(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}
