package ciphers

import (
	"strings"
)

type Trithemius struct{}

func (t *Trithemius) Encrypt(text string, key string) (string, error) {
	return encodeTrithemius(text), nil
}

func (t *Trithemius) Decrypt(text string, key string) (string, error) {
	return decodeTrithemius(text), nil
}

func (t *Trithemius) GetName() string {
	return "Шифр Тритемия"
}

func (t *Trithemius) GetDescription() string {
	return "Шифр с прогрессивным сдвигом, где каждая следующая буква сдвигается на одну позицию больше"
}

func (t *Trithemius) RequiresKey() bool {
	return false
}

func trithemius(text string, direction int) string {
	text = normalize(text)
	var result strings.Builder
	letterCount := 0

	for _, sim := range text {
		switch {
		case sim >= 'A' && sim <= 'Z':
			// берем letterCount == 0 т.к. все равно делать -1
			base := int('A')
			mod := 26
			currentIndex := int(sim) - base                          // 66 ('B') - 65 ('A')
			newIndex := (currentIndex + direction*letterCount) % mod // (1 + 1 * 0) % 26 = 1
			if newIndex < 0 {
				newIndex += mod
			}
			result.WriteRune(rune(base + newIndex)) // 65 + 1 = 66 => B
			letterCount++

		case sim >= 'a' && sim <= 'z':
			base := int('a')
			mod := 26
			currentIndex := int(sim) - base
			newIndex := (currentIndex + direction*letterCount) % mod
			if newIndex < 0 {
				newIndex += mod
			}
			result.WriteRune(rune(base + newIndex))
			letterCount++

		case sim >= 'А' && sim <= 'Я':
			base := int('А')
			mod := 32
			currentIndex := int(sim) - base
			newIndex := (currentIndex + direction*letterCount) % mod
			if newIndex < 0 {
				newIndex += mod
			}
			result.WriteRune(rune(base + newIndex))
			letterCount++

		case sim >= 'а' && sim <= 'я':
			base := int('а')
			mod := 32
			currentIndex := int(sim) - base
			newIndex := (currentIndex + direction*letterCount) % mod
			if newIndex < 0 {
				newIndex += mod
			}
			result.WriteRune(rune(base + newIndex))
			letterCount++

		default:
			result.WriteRune(sim)
		}
	}

	return result.String()
}

func encodeTrithemius(text string) string {
	return trithemius(text, 1)
}

func decodeTrithemius(text string) string {
	return trithemius(text, -1)
}

func init() {
	CiphersRegistry["Trithemius"] = &Trithemius{}
}
