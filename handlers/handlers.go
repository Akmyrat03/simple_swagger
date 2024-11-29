package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joefazee/api-doc/model"
)

type errResponse struct {
	Message string
}

type userResponse struct {
	Data []model.User `json:"data"`
}

type countryResponse struct {
	Data []model.Country `json:"data"`
}

// CreateUser creates a new user
// @Summary create a new user
// @Description create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} errResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
	}

	err := model.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, req)
}

// GetUsers return list of all users
// @Summary return list of all users
// @Description get all users
// @Tags Users
// @Success 200 {object} userResponse
// @Router /users/all [get]
func GetUsers(c *gin.Context) {
	users := model.ListUsers()

	c.JSON(http.StatusOK, userResponse{Data: users})
}

// GetCountries returns a list of all countries
// @Summary get all countries
// @Description get all countries
// @Tags Countries
// @Success 200 {object} countryResponse
// @Router /countries/all [get]
func GetCountries(c *gin.Context) {
	countries := model.ListCountries()

	c.JSON(http.StatusOK, countryResponse{Data: countries})
}
