package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	inputs := []string{
		"1234567",     // Standard integer
		"-12345.67",   // Negative float
		"+123456.789", // Positive sign float
		"12",          // Short integer
		"123.456",     // Short integer with float
		"-1234",       // Negative integer
	}

	for _, s := range inputs {
		fmt.Printf("Input: %-12s | Output: %s\n", s, comma(s))
	}
}

func comma(s string) string {
	var buf bytes.Buffer

	start := 0
	if len(s) > 0 && (s[0] == '+' || s[0] == '-') {
		buf.WriteByte(s[0])
		start = 1
	}
	end := strings.Index(s, ".")
	if end == -1 {
		end = len(s)
	}
	integerPart := s[start:end]
	n := len(integerPart)

	for i, c := range integerPart {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteRune(c)
	}

	if end < len(s) {
		buf.WriteString(s[end:])
	}

	return buf.String()
}
