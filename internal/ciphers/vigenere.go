package ciphers

import (
	"errors"
	"unicode"
)

type Vigenere struct {
	mode string // "plaintext" или "ciphertext"
}

func (v *Vigenere) Encrypt(text string, key string) (string, error) {
	if key == "" {
		return "", errors.New("ключ не может быть пустым")
	}
	// Преобразуем ключ в rune
	keyRune := []rune(key)
	if len(keyRune) == 0 {
		return "", errors.New("ключ не может быть пустым")
	}

	// В зависимости от режима выбираем функцию шифрования
	if v.mode == "ciphertext" {
		return vigenereAutokeyCiphertext(text, keyRune[0]), nil
	}
	// По умолчанию используем самоключ
	return vigenereAutokeyPlaintext(text, keyRune[0]), nil
}

func (v *Vigenere) Decrypt(text string, key string) (string, error) {
	if key == "" {
		return "", errors.New("ключ не может быть пустым")
	}
	// Преобразуем ключ в rune
	keyRune := []rune(key)
	if len(keyRune) == 0 {
		return "", errors.New("ключ не может быть пустым")
	}

	// В зависимости от режима выбираем функцию дешифрования
	if v.mode == "ciphertext" {
		return decryptVigenereAutokeyCiphertext(text, keyRune[0]), nil
	}
	// По умолчанию используем самоключ
	return decryptVigenereAutokeyPlaintext(text, keyRune[0]), nil
}

func (v *Vigenere) GetName() string {
	return "Шифр Виженера"
}

func (v *Vigenere) GetDescription() string {
	if v.mode == "ciphertext" {
		return "Шифр Виженера с использованием ключа-шифртекста"
	}
	return "Шифр Виженера с использованием самоключа (ключ - открытый текст)"
}

func (v *Vigenere) RequiresKey() bool {
	return true
}

func (v *Vigenere) SetMode(mode string) {
	v.mode = mode
}

// Функция для преобразования буквы в номер с учетом русского алфавита (без ё)
func letterToNumber(r rune) (int, int, bool) {
	// Английские буквы
	if r >= 'A' && r <= 'Z' {
		return int(r - 'A'), 26, true
	}
	if r >= 'a' && r <= 'z' {
		return int(r - 'a'), 26, true
	}

	// Русские буквы (без ё)
	if r >= 'А' && r <= 'Я' && r != 'Ё' {
		return int(r - 'А'), 32, true
	}
	if r >= 'а' && r <= 'я' && r != 'ё' {
		return int(r - 'а'), 32, true
	}

	return -1, 0, false // Не буква
}

// Функция для преобразования номера в букву с учетом алфавита
func numberToLetter(n int, upper bool, mod int) rune {
	n = n % mod
	if n < 0 {
		n += mod
	}

	if mod == 26 {
		if upper {
			return rune('A' + n)
		}
		return rune('a' + n)
	} else if mod == 32 {
		if upper {
			return rune('А' + n)
		}
		return rune('а' + n)
	}

	return rune(0)
}

// Шифр Виженера с самоключом (использует открытый текст)
func vigenereAutokeyPlaintext(text string, initialKey rune) string {
	var result []rune

	// Преобразуем текст в руны
	runes := []rune(text)

	// Обрабатываем первую букву с начальным ключом
	if len(runes) > 0 {
		prevKey, mod, isKeyLetter := letterToNumber(initialKey)
		if !isKeyLetter {
			prevKey = 0
			mod = 26 // По умолчанию используем английский алфавит
		}

		currentNumber, currentMod, isTextLetter := letterToNumber(runes[0])
		if isTextLetter {
			// Используем модуль текста, если ключ не подходит
			if mod != currentMod {
				mod = currentMod
			}

			newNumber := (currentNumber + prevKey) % mod
			isUpper := unicode.IsUpper(runes[0]) || (runes[0] >= 'А' && runes[0] <= 'Я')
			result = append(result, numberToLetter(newNumber, isUpper, mod))
			prevKey = currentNumber // Для следующей итерации
		} else {
			result = append(result, runes[0])
		}

		// Обрабатываем остальные буквы
		for i := 1; i < len(runes); i++ {
			currentNumber, currentMod, isTextLetter := letterToNumber(runes[i])
			if isTextLetter {
				// Используем модуль текста, если предыдущий ключ не подходит
				if mod != currentMod {
					mod = currentMod
				}

				newNumber := (currentNumber + prevKey) % mod
				isUpper := unicode.IsUpper(runes[i]) || (runes[i] >= 'А' && runes[i] <= 'Я')
				result = append(result, numberToLetter(newNumber, isUpper, mod))
				prevKey = currentNumber // Для следующей итерации
			} else {
				result = append(result, runes[i])
				// Для не-буквенных символов ключ не меняется
			}
		}
	}

	return string(result)
}

// Шифр Виженера с ключом-шифртекстом
func vigenereAutokeyCiphertext(text string, initialKey rune) string {
	var result []rune

	// Преобразуем текст в руны
	runes := []rune(text)

	// Обрабатываем первую букву с начальным ключом
	if len(runes) > 0 {
		prevKey, mod, isKeyLetter := letterToNumber(initialKey)
		if !isKeyLetter {
			prevKey = 0
			mod = 26 // По умолчанию используем английский алфавит
		}

		currentNumber, currentMod, isTextLetter := letterToNumber(runes[0])
		if isTextLetter {
			// Используем модуль текста, если ключ не подходит
			if mod != currentMod {
				mod = currentMod
			}

			newNumber := (currentNumber + prevKey) % mod
			isUpper := unicode.IsUpper(runes[0]) || (runes[0] >= 'А' && runes[0] <= 'Я')
			result = append(result, numberToLetter(newNumber, isUpper, mod))
			prevKey = newNumber // Для следующей итерации
		} else {
			result = append(result, runes[0])
		}

		// Обрабатываем остальные буквы
		for i := 1; i < len(runes); i++ {
			currentNumber, currentMod, isTextLetter := letterToNumber(runes[i])
			if isTextLetter {
				// Используем модуль текста, если предыдущий ключ не подходит
				if mod != currentMod {
					mod = currentMod
				}

				newNumber := (currentNumber + prevKey) % mod
				isUpper := unicode.IsUpper(runes[i]) || (runes[i] >= 'А' && runes[i] <= 'Я')
				result = append(result, numberToLetter(newNumber, isUpper, mod))
				prevKey = newNumber // Для следующей итерации
			} else {
				result = append(result, runes[i])
				// Для не-буквенных символов ключ не меняется
			}
		}
	}

	return string(result)
}

// Дешифрование для шифра с самоключом
func decryptVigenereAutokeyPlaintext(ciphertext string, initialKey rune) string {
	var result []rune

	// Преобразуем текст в руны
	runes := []rune(ciphertext)

	// Обрабатываем первую букву с начальным ключом
	if len(runes) > 0 {
		prevKey, mod, isKeyLetter := letterToNumber(initialKey)
		if !isKeyLetter {
			prevKey = 0
			mod = 26 // По умолчанию используем английский алфавит
		}

		currentNumber, currentMod, isTextLetter := letterToNumber(runes[0])
		if isTextLetter {
			// Используем модуль текста, если ключ не подходит
			if mod != currentMod {
				mod = currentMod
			}

			newNumber := (currentNumber - prevKey) % mod
			if newNumber < 0 {
				newNumber += mod
			}
			isUpper := unicode.IsUpper(runes[0]) || (runes[0] >= 'А' && runes[0] <= 'Я')
			result = append(result, numberToLetter(newNumber, isUpper, mod))
			prevKey = newNumber // Для следующей итерации (используем расшифрованную букву)
		} else {
			result = append(result, runes[0])
		}

		// Обрабатываем остальные буквы
		for i := 1; i < len(runes); i++ {
			currentNumber, currentMod, isTextLetter := letterToNumber(runes[i])
			if isTextLetter {
				// Используем модуль текста, если предыдущий ключ не подходит
				if mod != currentMod {
					mod = currentMod
				}

				newNumber := (currentNumber - prevKey) % mod
				if newNumber < 0 {
					newNumber += mod
				}
				isUpper := unicode.IsUpper(runes[i]) || (runes[i] >= 'А' && runes[i] <= 'Я')
				result = append(result, numberToLetter(newNumber, isUpper, mod))
				prevKey = newNumber // Для следующей итерации
			} else {
				result = append(result, runes[i])
				// Для не-буквенных символов ключ не меняется
			}
		}
	}

	return string(result)
}

// Дешифрование для шифра с ключом-шифртекстом
func decryptVigenereAutokeyCiphertext(ciphertext string, initialKey rune) string {
	var result []rune

	// Преобразуем текст в руны
	runes := []rune(ciphertext)

	// Обрабатываем первую букву с начальным ключом
	if len(runes) > 0 {
		prevKey, mod, isKeyLetter := letterToNumber(initialKey)
		if !isKeyLetter {
			prevKey = 0
			mod = 26 // По умолчанию используем английский алфавит
		}

		currentNumber, currentMod, isTextLetter := letterToNumber(runes[0])
		if isTextLetter {
			// Используем модуль текста, если ключ не подходит
			if mod != currentMod {
				mod = currentMod
			}

			newNumber := (currentNumber - prevKey) % mod
			if newNumber < 0 {
				newNumber += mod
			}
			isUpper := unicode.IsUpper(runes[0]) || (runes[0] >= 'А' && runes[0] <= 'Я')
			result = append(result, numberToLetter(newNumber, isUpper, mod))
			prevKey = currentNumber // Для следующей итерации (используем букву шифртекста)
		} else {
			result = append(result, runes[0])
		}

		// Обрабатываем остальные буквы
		for i := 1; i < len(runes); i++ {
			currentNumber, currentMod, isTextLetter := letterToNumber(runes[i])
			if isTextLetter {
				// Используем модуль текста, если предыдущий ключ не подходит
				if mod != currentMod {
					mod = currentMod
				}

				newNumber := (currentNumber - prevKey) % mod
				if newNumber < 0 {
					newNumber += mod
				}
				isUpper := unicode.IsUpper(runes[i]) || (runes[i] >= 'А' && runes[i] <= 'Я')
				result = append(result, numberToLetter(newNumber, isUpper, mod))
				prevKey = currentNumber // Для следующей итерации
			} else {
				result = append(result, runes[i])
				// Для не-буквенных символов ключ не меняется
			}
		}
	}

	return string(result)
}

func init() {
	CiphersRegistry["Vigenere"] = &Vigenere{mode: "plaintext"}
}
