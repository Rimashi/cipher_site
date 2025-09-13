package models

type Cipher struct {
	ID          string
	Name        string
	Description string
	HasKey      bool
	KeyLabel    string // Например: "Сдвиг", "Ключевое слово"
}
