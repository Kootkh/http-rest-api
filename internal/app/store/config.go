package store

// ------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------

// Config - структура конфигурации хранилища
type Config struct {

	// Строка подключения к базе данных
	DatabaseURL string `toml:"database_url"`
}

// ------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------

// NewConfig - конструктор нового инстанса Config
//
// Принимает: config - указатель на структуру Config которая содержит конфигурацию
//
// Возвращает: указатель на новый инстанс Store
func NewConfig() *Config {

	return &Config{}

}
