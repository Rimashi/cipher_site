package ciphers

import (
	"strings"
)

type Atbash struct{}

func (a *Atbash) Encrypt(text string, key string) (string, error) {
	return atbashEncode(text), nil
}

func (a *Atbash) Decrypt(text string, key string) (string, error) {
	return atbashEncode(text), nil
}

func (a *Atbash) GetName() string {
	return "Атбаш"
}

func (a *Atbash) GetDescription() string {
	return "Шифр замены, где каждая буква алфавита заменяется на зеркальную"
}

func (a *Atbash) RequiresKey() bool {
	return false
}

func atbashEncode(text string) string {
	text = normalize(text)
	var result strings.Builder

	for _, sim := range text {
		switch {
		case sim >= 'A' && sim <= 'Z':
			newSim := 'Z' - (sim - 'A')
			result.WriteRune(newSim)
		case sim >= 'a' && sim <= 'z':
			newSim := 'z' - (sim - 'a')
			result.WriteRune(newSim)
		case sim >= 'А' && sim <= 'Я':
			newSim := 'Я' - (sim - 'А')
			result.WriteRune(newSim)
		case sim >= 'а' && sim <= 'я':
			newSim := 'я' - (sim - 'а')
			result.WriteRune(newSim)
		default:
			result.WriteRune(sim)
		}
	}

	return result.String()
}

func init() {
	CiphersRegistry["Atbash"] = &Atbash{}
}
