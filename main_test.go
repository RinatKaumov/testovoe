package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerValidRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=2", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался код 200")

	// Проверяем тело ответа
	expectedBody := strings.Join(cafeList["moscow"][:2], ",")
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ не соответствует ожидаемому")
}

func TestMainHandlerInvalidCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=unknown&count=2", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался код 400")

	// Проверяем тело ответа
	expectedBody := "wrong city value"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ не соответствует ожидаемому")
}

func TestMainHandlerMissingCount(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code, "Ожидался код 400")

	// Проверяем тело ответа
	expectedBody := "count missing"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ не соответствует ожидаемому")
}

func TestMainHandlerCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=10", nil)
	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// Проверяем код ответа
	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Ожидался код 200")

	// Проверяем тело ответа
	expectedBody := strings.Join(cafeList["moscow"], ",")
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Ответ не соответствует ожидаемому")
}
