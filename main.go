package mypackage

import (
	"fmt"

	"github.com/bosshie96/sayhi/utils" // Importing the utils package
)

// SayHi prints a simple greeting message.
func SayHi(name string) {
	fmt.Println("Hi %s, This is my first package module in Go!", name)
}

// FormatName converts a given string to uppercase letters.
func FormatName(name string) string {
	return utils.FormatName(name)
}
