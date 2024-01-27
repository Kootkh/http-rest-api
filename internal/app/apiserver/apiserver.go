package apiserver

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Описание структуры APIServer ...
type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// NewAPIServer  - конфигураторе, создание APIServer..
//
// Возвращает указатель на новый сконфигурировнный инстанс APIServer.
func NewAPIServer(config *Config) *APIServer {

	// Инициализируем APIServer
	return &APIServer{

		// Указываем что в атрибуте config структуры APIServer ссылаемся на параметр конфигурации config
		config: config,

		// Инициализируем новый инстанс логгера
		logger: logrus.New(),

		// Инициализируем новый инстанс маршрутизатора
		router: mux.NewRouter(),
	}

}

// Start - запуск APIServer
//
// Возвращает ошибку в случае возникновения
func (s *APIServer) Start() error {

	// Вызываем конфигуратор инстанса логгера и проверяем на ошибки
	if err := s.configureLogger(); err != nil {

		// Если произошла ошибка, то возвращаем её.
		return err

	}

	// Вызываем конфигуратор инстанса маршрутизатора
	s.configureRouter()

	// Поскольку конфигуратор инстанса маршрутизатора ошибок не возвращает, то мы их не проверяем.

	// Если на предыдущих этапах ошибок не возникло, то выводим в лог информацию о запуске APIServer (уровень INFO)
	s.logger.Info("starting api server")

	// Вызываем метод ListenAndServe из пакета http с указанием параметра BindAddr из нашего инстанса конфигурации и полученного ранее инстанса маршрутизатора router.
	// Функция http.ListenAndServe всегда возвращает non-nil ошибку, которую мы и будем возвращать в нашей функции Start.
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// configureLogger - конфигуратор нового инстанса логгера
func (s *APIServer) configureLogger() error {

	// С помощью функции ParseLevel из пакета logrus парсим уровень логирования из инстанса конфигурации config
	level, err := logrus.ParseLevel(s.config.LogLevel)

	// Проверяем на ошибки
	if err != nil {

		// Если произошла ошибка, то возвращаем её.
		return err

	}

	// Если ошибок нет, то ставим нашему логгеру соответствующий уровень
	s.logger.SetLevel(level)

	// Выходим из функции и возвращаем nil
	return nil

}

// configureRouter - конфигуратор нового инстанса маршрутизатора
func (s *APIServer) configureRouter() {

	// Добавляем тестовый маршрут
	s.router.HandleFunc("/hello", s.handleHello())

}

// handleHello - обработчик маршрута /hello
//
// Возвращает функцию-обработчик http.HandlerFunc
func (s *APIServer) handleHello() http.HandlerFunc {

	// Тут можно определить:
	// 	- какие-то переменные, которые будут использоваться только в этом обработчике маршрутов, и этот код выполнится всего один раз.
	//	- какие-то локальные специфичные типы, например:
	// 		type request struct {	// Описание типа структуры входящего запроса
	//			name string
	//			...
	// 		}
	// 		type response struct {	// Описание типа структуры возвращаемого ответа
	//			...
	// 		}

	// Возвращаем функцию http.HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {

		// Тут описывается вся бизнес-логика обработчика маршрута и логика обработки каждого конкретного запроса.

		// Для примера:
		io.WriteString(w, "Hello")

	}
}
