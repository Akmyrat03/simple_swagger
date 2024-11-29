package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joefazee/api-doc/model"
)

type errResponse struct {
	Message string
}

// CreateUser creates a new user
// @Summary create a new user
// @Description create new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body model.User true "User"
// @Success 201 {object} model.User
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
