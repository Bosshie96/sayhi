package utils

import "strings"

func FormatName(name string) string {
	return strings.ToUpper(name) + "!\n"
}
