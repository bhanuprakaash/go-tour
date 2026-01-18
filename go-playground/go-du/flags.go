package main

import (
	"fmt"
	"strings"
)

type Exclude []string

func (e *Exclude) Set(value string) error {
	if value == "" {
		return nil
	}
	parts := strings.Split(value, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			*e = append(*e, part)
		}
	}
	return nil
}
func (e *Exclude) String() string {
	return fmt.Sprintf("%v", *e)
}
