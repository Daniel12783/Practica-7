package controller

import (
	"encoding/json"
	"gin_http/cmd/services"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (s *UserController) GetUsers(c *gin.Context) {
	users := s.userService.GetUsers()
	c.JSON(http.StatusOK, users)
}

func (s *UserController) CreateUser(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsin body",
		})
		return
	}

	var user services.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing body",
		})
		return
	}
	user = s.userService.CreateUser(user)

	c.JSON(http.StatusOK, user)
}

func (s *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var updateUser services.User
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	//Convert id to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	user, err := s.userService.UpdateUser(idInt, updateUser)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	err = s.userService.DeleteUser(idInt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
		"id":      id,
	})
}
