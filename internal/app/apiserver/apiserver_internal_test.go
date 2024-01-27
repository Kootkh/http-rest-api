package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_HandleHello(t *testing.T) {

	// Создаем новый инстанс сервера APIServer, который принимает на вход конфигурацию от функции NewConfig()
	s := NewAPIServer(NewConfig())

	// С помощью метода NewRecorder создаём новый инстанс ResponseRecorder-а "rec", в котором будем хранить информацию о ответе сервера
	rec := httptest.NewRecorder()

	// С помощью метода NewRequest создаём новый инстанс запроса "req", в который передаём метод "GET", путь "/hello" и "nil" в качестве содержания тела запроса.
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	// Вызываем обработчик маршрута handleHello() сервера APIServer, в который передаем наши инстансы "rec" и "req"
	s.handleHello().ServeHTTP(rec, req)

	// С помощью метода Equal из библиотеки github.com/stretchr/testify/assert проверяем что содержимое тела ответа "rec" сервера равно тому что должно возвращается нашим обработчиком ("func handleHello()" -> apiserver.go)
	assert.Equal(t, rec.Body.String(), "Hello")
}
