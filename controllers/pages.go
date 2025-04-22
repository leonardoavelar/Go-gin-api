package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PageNotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
