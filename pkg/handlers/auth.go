package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SignIn",
	})
}

func SignUp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SignUp",
	})
}
