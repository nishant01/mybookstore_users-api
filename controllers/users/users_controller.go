package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nishant01/mybookstore_users-api/doamin/users"
	"github.com/nishant01/mybookstore_users-api/services"
	"github.com/nishant01/mybookstore_users-api/utils/errors"
)

// CreateUser User
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

// UpdateUser User
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Update User")
}

// DeleteUser User
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement Delete User")
}

// GetUser User
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUserById(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

// func ListUser(c *gin.Context) {
// 	c.String(http.StatusOK, "Implement List User")
// }

// func SearchUser(c *gin.Context) {
// 	c.String(http.StatusOK, "Implement Search User")
// }
