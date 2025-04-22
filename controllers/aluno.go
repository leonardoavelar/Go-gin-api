package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardoavelar/Go-gin-api/database"
	"github.com/leonardoavelar/Go-gin-api/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// GetAlunos 	godoc
// @Sumary 		Consulta todos os alunos cadastrados
// @Description get json
// @Tags 		Aluno
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} []models.Aluno
// @Failure 	400 {object} httputil.HTTPError
// @Failure 	404 {object} httputil.HTTPError
// @Failure 	500 {object} httputil.HTTPError
// @Router 		/aluno [get]
func GetAlunos(c *gin.Context) {

	var alunos []models.Aluno

	database.DB.Find(&alunos)

	c.JSON(http.StatusOK, gin.H{
		"data": alunos})
}

// GetAlunoById	godoc
// @Sumary 		Consulta um  aluno por Id
// @Description get json
// @Tags 		Aluno
// @Accept 		json
// @Produce 	json
// @Param		id 	path 	 int true "Aluno ID"
// @Success 	200 {object} models.Aluno
// @Failure 	400 {object} httputil.HTTPError
// @Failure 	404 {object} httputil.HTTPError
// @Failure 	500 {object} httputil.HTTPError
// @Router 		/aluno/{id} [get]
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

// GetAlunoByName
// @Sumary 			Consulta um aluno por Nome
// @Description 	get json
// @Tags 			Aluno
// @Accept 			json
// @Produce 		json
// @Param			nome path 	 string true "Aluno Nome"
// @Success 		200 {object} models.Aluno
// @Failure 		400 {object} httputil.HTTPError
// @Failure 		404 {object} httputil.HTTPError
// @Failure 		500 {object} httputil.HTTPError
// @Router 			/aluno/nome/{nome} [get]
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

// PostAluno 	godoc
// @Sumary 		Cadastra novo aluno
// @Description post json
// @Tags 		Aluno
// @Accept 		json
// @Produce 	json
// @Param		aluno body 	 models.Aluno true "Aluno"
// @Success 	200 {object} models.Aluno
// @Failure 	400 {object} httputil.HTTPError
// @Failure 	500 {object} httputil.HTTPError
// @Router 		/aluno [post]
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

// DeleteAlunoById	godoc
// @Sumary 			Exclui aluno por Id
// @Description 	delete json
// @Tags 			Aluno
// @Accept 			json
// @Produce 		json
// @Param			id 	path 	 int true "Aluno ID"
// @Success 		200 {object} string
// @Failure 		400 {object} httputil.HTTPError
// @Failure 		500 {object} httputil.HTTPError
// @Router 			/aluno/{id} [delete]
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
