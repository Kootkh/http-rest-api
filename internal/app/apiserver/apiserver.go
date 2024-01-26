package apiserver

// Описание структуры APIServer ...
type APIServer struct {
}

// NewAPIServer  - конфигураторе, создание APIServer..
//
// Возвращает указатель на новый сконфигурировнный инстанс APIServer.
func NewAPIServer() *APIServer {
	// Инициализируем APIServer
	return &APIServer{}
}

// Start - запуск APIServer
//
// Возвращает ошибку в случае возникновения
func (s *APIServer) Start() error {
	return nil
}
