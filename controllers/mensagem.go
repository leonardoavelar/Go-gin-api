package controllers

import "github.com/gin-gonic/gin"

func Saudacao(c *gin.Context) {

	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza?",
	})
}
