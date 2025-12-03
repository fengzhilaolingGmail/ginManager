// utils/ptr.go
package utils

func StringPtrVal(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
