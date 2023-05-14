package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WeslleySanto/go-validacoes-testes-paginas-html/controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func Test_ShouldStatusCodeWithParams(t *testing.T) {
	r := SetupRotasTeste()

	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/weslley", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code)
	mockResponse := `{"API diz:":"E ai weslley, tudo beleza?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}
