package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoavelar/Go-gin-api/database"
	"github.com/leonardoavelar/Go-gin-api/models"
)

func GetAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(http.StatusOK, gin.H{
		"data": alunos})
}

func GetAlunoById(c *gin.Context) {

	id := c.Params.ByName("id")

	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Aluno não encontrado"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": aluno})
}

func GetAlunoByName(c *gin.Context) {

	nome := c.Params.ByName("nome")

	var aluno models.Aluno
	database.DB.Where(models.Aluno{Nome: nome}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Aluno não encontrado"})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": aluno})
}

func PostAluno(c *gin.Context) {

	var aluno models.Aluno
	err := c.ShouldBindJSON(&aluno)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

		return
	}

	validate := models.AlunoValidate(&aluno)

	if validate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validate.Error()})

		return
	}

	database.DB.Create(&aluno)

	c.JSON(http.StatusOK, aluno)
}

func DeleteAlunoById(c *gin.Context) {

	id := c.Params.ByName("id")

	var aluno models.Aluno
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Aluno não encontrado"})
}

func PatchAluno(c *gin.Context) {

	id := c.Params.ByName("id")

	var aluno models.Aluno
	database.DB.First(&aluno, id)

	err := c.ShouldBindJSON(&aluno)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})

		return
	}

	validate := models.AlunoValidate(&aluno)

	if validate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": validate.Error()})

		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

func AlunoPageIndex(c *gin.Context) {

	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})

}
