package myStr

func ToUpper(s string) string {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			continue
		} else if b[i] >= 'a' && b[i] <= 'z' {
			b[i] = b[i] - 32
		}
	}
	return string(b)
}
func ToLower(s string) string {
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] = b[i] + 32
		} else if b[i] >= 'a' && b[i] <= 'z' {
			continue
		}
	}
	return string(b)
}
func Title(s string) string {
	if s[0] >= 'A' && s[0] <= 'Z' {
		return s
	}
	b := []byte(s)
	if b[0] >= 'a' && b[0] <= 'z' {
		b[0] = b[0] - 32
	}
	return string(b)
}
func Count(s, h string) int {
	count := 0
	t := []byte(h)
	for i := 0; i < len(s); i++ {
		if s[i] == t[0] {
			count++
		}
	}
	return count
}
