package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUpHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func SignInHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}
