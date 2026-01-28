package ciphers

import (
	"errors"
	"strings"
)

type Belazo struct{}

func (b *Belazo) Encrypt(text string, key string) (string, error) {
	if key == "" {
		return "", errors.New("ключ не может быть пустым")
	}
	return encryptBelazo(text, key), nil
}

func (b *Belazo) Decrypt(text string, key string) (string, error) {
	if key == "" {
		return "", errors.New("ключ не может быть пустым")
	}
	return decryptBelazo(text, key), nil
}

func (b *Belazo) GetName() string {
	return "Шифр Белазо"
}

func (b *Belazo) GetDescription() string {
	return "Шифр на основе гаммирования с использованием ключевого слова"
}

func (b *Belazo) RequiresKey() bool {
	return true
}

// Определяем позицию буквы в алфавите и возвращаем (позиция, мощность_алфавита, является_ли_буквой)
func getLetterPosition(r rune) (int, int, bool) {
	switch {
	case r >= 'A' && r <= 'Z': // Английские заглавные
		return int(r - 'A'), 26, true
	case r >= 'a' && r <= 'z': // Английские строчные
		return int(r - 'a'), 26, true
	case r >= 'А' && r <= 'Я': // Русские заглавные
		return int(r - 'А'), 32, true
	case r >= 'а' && r <= 'я': // Русские строчные
		return int(r - 'а'), 32, true
	default: // Не буква
		return 0, 0, false
	}
}

// Получаем букву по позиции в алфавите
func getLetterFromPosition(pos int, isUpper bool, mod int) rune {
	// Корректируем отрицательные позиции
	if pos < 0 {
		pos += mod
	}
	// Обеспечиваем циклический сдвиг
	pos %= mod

	// Возвращаем букву в нужном регистре
	switch mod {
	case 26: // Английский алфавит
		if isUpper {
			return rune('A' + pos)
		}
		return rune('a' + pos)
	case 32: // Русский алфавит
		if isUpper {
			return rune('А' + pos)
		}
		return rune('а' + pos)
	default: // Неизвестный алфавит
		return rune(0)
	}
}

// Основная функция шифрования/дешифрования
func belazo(text, key string, encrypt bool) string {
	// Нормализуем текст и ключ
	text = normalize(text)
	key = normalize(key)

	var result strings.Builder // Для построения результата

	// Создаем расширенный ключ той же длины, что и текст
	textRunes := []rune(text)
	keyRunes := []rune(key)
	keyLen := len(keyRunes)
	keyIndex := 0

	for _, sim := range textRunes {

		textPos, mod, isTextLetter := getLetterPosition(sim)
		if isTextLetter {

			keyRune := keyRunes[keyIndex%keyLen]
			keyPos, _, _ := getLetterPosition(keyRune)

			isUpper := (sim >= 'A' && sim <= 'Z') || (sim >= 'А' && sim <= 'Я')

			var newPos int
			if encrypt {
				newPos = (textPos + keyPos) % mod
			} else {
				newPos = (textPos - keyPos) % mod
			}

			result.WriteRune(getLetterFromPosition(newPos, isUpper, mod))
			keyIndex++ // 🔥 ключ сдвигается ТОЛЬКО на буквах
			continue
		}

		// пробелы и знаки — как есть
		result.WriteRune(sim)
	}

	return result.String() // Возвращаем результат
}

// Функции-обертки для удобства
func encryptBelazo(text, key string) string {
	return belazo(text, key, true) // Шифрование
}

func decryptBelazo(text, key string) string {
	return belazo(text, key, false) // Дешифрование
}

func init() {
	CiphersRegistry["Belazo"] = &Belazo{}
}
