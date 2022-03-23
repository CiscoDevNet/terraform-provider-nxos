package helpers

func ParseNxosBoolean(s string) bool {
	if s == "yes" || s == "true" {
		return true
	}
	return false
}
