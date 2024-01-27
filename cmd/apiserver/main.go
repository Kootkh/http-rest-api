package main

// https://www.youtube.com/playlist?list=PLehOyJfJkFkJ5m37b4oWh783yzVlHdnUH

import (
	"flag"
	"http-rest-api/internal/app/apiserver"
	"log"

	"github.com/BurntSushi/toml"
)

// Что бы мы могли задавать путь к файлу конфига в качестве флага для нашего бинарника
var (
	configPath string
)

// init - инициализация
func init() {
	// Парсим флаги
	// flag.StringVar(в какую переменную парсим, первый аргумент - имя флага, второй аргумент - значение по умолчанию, третий аргумент - описание)
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

// main - точка входа
func main() {

	// Запускаем парсинг переданых при запуске флагов
	flag.Parse()

	// С помощью функции apiserver.NewConfig создаём новый инстанс конфига для функции NewAPIServer
	config := apiserver.NewConfig()

	// Декодируем конфигурационный файл с помощью функции toml.DecodeFile из пакета github.com/BurntSushi/toml
	// toml.DecodeFile(путь к конфигурационному файлу, инстанс конфига в который будем сохранять конфигурацию)
	_, err := toml.DecodeFile(configPath, config)

	// Проверяем на ошибки
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем APIServer с передачей нашего инстанса конфига в качестве аргумента
	server := apiserver.NewAPIServer(config)

	// Запускаем APIServer
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
