package mypackage

import (
	"fmt"
	"strings"
)

// SayHi prints a simple greeting message.
func SayHi(name string) {
	fmt.Println("Hi %s, This is my first package module in Go!", name)
}

// FormatName converts a given string to uppercase letters.
func FormatName(name string) string {
	return strings.ToUpper(name) + "!\n"
}
