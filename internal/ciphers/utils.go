package ciphers

import "strings"

func normalize(text string) string {
	text = strings.ReplaceAll(text, "ё", "е")
	text = strings.ReplaceAll(text, "Ё", "Е")
	return text
}
