package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAuthHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}
