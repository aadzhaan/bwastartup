package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	//map input dari user ke struct RegisterUserInput
	//struct di atas kita passing sebagai parameter service
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		//gin.H adalah mapping dimana gin merupakan string, value interface (bisa apa aja)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Failed to Registered", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIResponse("Failed to Registered", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// token, err := h.jwtService.GenerateToken()

	formatter := user.FormatUser(newUser)
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	//user input email dan password
	//input di tangkap handler
	//mapping dari input user ke input struct
	//input struct passing service
	//di service mencari dg bantuan repository user dengan email x
	//mencocokkan password

	//proses untuk check validasi input dari data yang dimasukkan
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		//gin.H adalah mapping dimana gin merupakan string, value interface (bisa apa aja)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//cek ke service(jquery function)
	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to Login", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser)
	response := helper.APIResponse("Login successfully", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
