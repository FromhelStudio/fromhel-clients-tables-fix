package services

import (
	"fmt"
	"strings"
)

func Trim(text interface{}) (string, error) {
	switch v := text.(type) {
	case string:
		return strings.Join(strings.Fields(strings.TrimSpace(v)), " "), nil
	default:
		return "", fmt.Errorf("invalid string")
	}
}
