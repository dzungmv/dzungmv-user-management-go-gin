package utils

import "strings"

func NomarlizeString(text string) string {

	return strings.ToLower(strings.TrimSpace(text))

}
