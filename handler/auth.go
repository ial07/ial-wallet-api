package handler

import (
	"fmt"
	"ial-wallet/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)
	

func LoginHandler(c *gin.Context) {
	var data model.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
	}
	
	c.JSON(http.StatusOK, gin.H{
		"email":    data.Email,
		"password": data.Password,
	})

}

func RegisterHandler(c *gin.Context) {
	var data model.User

	err := c.ShouldBindJSON(&data)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
	}
	

	c.JSON(http.StatusOK, gin.H{
		"email":    data.Email,
		"password": data.Password,
		"nohp":   data.Nohp,
	})
}