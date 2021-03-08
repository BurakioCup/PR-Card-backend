package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReadCardHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func ReadCardsHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func ReadMycardHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func CreateCard()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func UpdateCard()gin.HandlerFunc{
	return func(c *gin.Context) {

	c.JSON(http.StatusOK, "")
	}
}
