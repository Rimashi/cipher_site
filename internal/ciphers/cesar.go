package ciphers

import (
	"errors"
	"strconv"
	"strings"
)

type Cesar struct{}

func (c *Cesar) Encrypt(text string, key string) (string, error) {
	shift, err := strconv.Atoi(key)
	if err != nil {
		return "", errors.New("ключ должен быть числом")
	}
	return cesarEncrypt(text, shift), nil
}

func (c *Cesar) Decrypt(text string, key string) (string, error) {
	shift, err := strconv.Atoi(key)
	if err != nil {
		return "", errors.New("ключ должен быть числом")
	}
	return cesarDecrypt(text, shift), nil
}

func (c *Cesar) GetName() string {
	return "Шифр Цезаря"
}

func (c *Cesar) GetDescription() string {
	return "Шифр сдвига, где каждая буква алфавита сдвигается на фиксированное число позиций"
}

func (c *Cesar) RequiresKey() bool {
	return true
}

func cesarEncrypt(text string, shift int) string {
	return cesar(text, shift)
}

func cesarDecrypt(text string, shift int) string {
	return cesar(text, -shift)
}

func cesar(text string, shift int) string {
	text = normalize(text)
	var result strings.Builder

	for _, sim := range text {
		switch {
		case sim >= 'A' && sim <= 'Z':
			base := int('A')
			mod := 26
			pos := (int(sim) - base + shift) % mod
			if pos < 0 {
				pos += mod
			}
			result.WriteRune(rune(base + pos))

		case sim >= 'a' && sim <= 'z':
			base := int('a')
			mod := 26
			pos := (int(sim) - base + shift) % mod
			if pos < 0 {
				pos += mod
			}
			result.WriteRune(rune(base + pos))

		case sim >= 'А' && sim <= 'Я':
			base := int('А')
			mod := 32
			pos := (int(sim) - base + shift) % mod
			if pos < 0 {
				pos += mod
			}
			result.WriteRune(rune(base + pos))

		case sim >= 'а' && sim <= 'я':
			base := int('а')
			mod := 32
			pos := (int(sim) - base + shift) % mod
			if pos < 0 {
				pos += mod
			}
			result.WriteRune(rune(base + pos))
		default:
			result.WriteRune(sim)
		}
	}

	return result.String()
}

func init() {
	CiphersRegistry["Cesar"] = &Cesar{}
}
