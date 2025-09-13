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

// Создаем расширенный ключ нужной длины
func expandKey(key string, length int) []rune {
	// Преобразуем строку ключа в срез рун
	keyRunes := []rune(key)

	// Создаем новый срез рун нужной длины
	expandedKey := make([]rune, length)

	// Заполняем расширенный ключ, циклически повторяя исходный ключ
	for i := 0; i < length; i++ {
		// Берем символ из ключа по модулю его длины
		// Это обеспечивает циклическое повторение ключа
		expandedKey[i] = keyRunes[i%len(keyRunes)]
	}

	// Возвращаем расширенный ключ
	return expandedKey
}

// Основная функция шифрования/дешифрования
func belazo(text, key string, encrypt bool) string {
	// Нормализуем текст и ключ
	text = normalize(text)
	key = normalize(key)

	var result strings.Builder // Для построения результата

	// Создаем расширенный ключ той же длины, что и текст
	textRunes := []rune(text)
	expandedKey := expandKey(key, len(textRunes))

	letterCount := 0 // Счетчик букв (игнорируем не-буквы)

	// Обрабатываем каждый символ текста
	for i, sim := range textRunes {
		// Получаем позицию и мощность алфавита для символа ключа
		if keyPos, mod, isKeyLetter := getLetterPosition(expandedKey[i]); isKeyLetter {
			// Получаем позицию и мощность алфавита для символа текста
			if textPos, _, isTextLetter := getLetterPosition(sim); isTextLetter {
				// Определяем регистр буквы
				isUpper := (sim >= 'A' && sim <= 'Z') || (sim >= 'А' && sim <= 'Я')

				var newPos int
				if encrypt {
					// Шифрование: добавляем позицию ключа
					newPos = (textPos + keyPos) % mod
				} else {
					// Дешифрование: вычитаем позицию ключа
					newPos = (textPos - keyPos) % mod
				}

				// Получаем новую букву и добавляем в результат
				result.WriteRune(getLetterFromPosition(newPos, isUpper, mod))
				letterCount++
				continue // Переходим к следующему символу
			}
		}
		// Если символ не буква - добавляем без изменений
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
