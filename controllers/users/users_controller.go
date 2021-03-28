package users

import (
	"github.com/gin-gonic/gin"
	"movies-user/domain/users"
	"movies-user/utils/rest_errors"
	"movies-user/validators"
	"net/http"
)

func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil{
		restErr := rest_errors.NewBadRequestError("invalid json Body")
		c.JSON(restErr.Status(),restErr)
		return
	}
	validateErr := validators.Validate(&user)
	if validateErr != nil{
		c.JSON(validateErr.Status(),validateErr)
		return
	}
	c.JSON(http.StatusCreated, user)
}
