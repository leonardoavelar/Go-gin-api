package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardoavelar/Go-gin-api/controllers"
)

func HandlerRequests() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Saudação
	r.GET("/:nome", controllers.Saudacao)

	// Aluno
	r.GET("/aluno", controllers.GetAlunos)
	r.GET("/aluno/:id", controllers.GetAlunoById)
	r.GET("/aluno/nome/:nome", controllers.GetAlunoByName)
	r.POST("/aluno", controllers.PostAluno)
	r.DELETE("/aluno/:id", controllers.DeleteAlunoById)
	r.PATCH("/aluno/:id", controllers.PatchAluno)

	// Pages
	r.GET("/aluno/index", controllers.AlunoPageIndex)

	r.NoRoute(controllers.PageNotFound)
	r.Run(":8000")
}
