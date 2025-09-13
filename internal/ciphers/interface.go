package ciphers

// Cipher интерфейс для всех шифров
type Cipher interface {
	Encrypt(text string, key string) (string, error)
	Decrypt(text string, key string) (string, error)
	GetName() string
	GetDescription() string
	RequiresKey() bool
}

// Регистр всех доступных шифров
var CiphersRegistry = map[string]Cipher{}
