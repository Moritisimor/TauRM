package helpers

import "strings"

func IsIgnorable(line string) bool {
	return strings.HasSuffix(line, ":") || strings.TrimSpace(line) == "" || strings.HasPrefix(line ,"#")
}
