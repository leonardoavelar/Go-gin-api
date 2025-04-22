package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leonardoavelar/Go-gin-api/controllers"
	"github.com/leonardoavelar/Go-gin-api/database"
	"github.com/stretchr/testify/assert"
)

func SetupRoutesTest() *gin.Engine {

	routes := gin.Default()
	return routes

}

func TestSaudacaoOk(t *testing.T) {

	r := SetupRoutesTest()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/Leonardo", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Devem ser iguais")

	mockResp := `{"API diz":"E ai Leonardo, tudo beleza?"}`
	respBody, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, mockResp, string(respBody))
}

func TestGetAlunosOK(t *testing.T) {

	database.Conectar()

	r := SetupRoutesTest()
	r.GET("/aluno", controllers.GetAlunos)

	req, _ := http.NewRequest("GET", "/aluno", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetAlunoByNameOK(t *testing.T) {

	database.Conectar()

	r := SetupRoutesTest()
	r.GET("/aluno/nome/:nome", controllers.GetAlunoByName)

	req, _ := http.NewRequest("GET", "/aluno/nome/Leonardo Avelar", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
