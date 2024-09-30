package handlers

import (
	"mows-game-center-time-mgt/services"
	"mows-game-center-time-mgt/utils/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get all users
func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Could't retrieve users", nil, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "Successfully retrived all users", users, nil))
}

func GetUserByID(c *gin.Context){
	id , err := strconv.Atoi(c.Param("id"))
	if err != nil{
		errRes := response.ClientResponse(http.StatusBadRequest,"Invalid user ID",nil,err.Error())
		c.JSON(http.StatusBadRequest,errRes)
		return
	}
	user , err := services.GetUserByID(id)
	if err != nil{
		errRes := response.ClientResponse(http.StatusNotFound,"User not found",nil,err.Error())
		c.JSON(http.StatusNotFound,errRes)
	}
	successRes := response.ClientResponse(http.StatusOK,"User retrived successfully",user,err.Error())
	c.JSON(http.StatusOK,successRes)
}
// Delete User
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ClientResponse(http.StatusBadRequest, "Invalid user ID", nil, err.Error()))
		return
	}

	err = services.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ClientResponse(http.StatusInternalServerError, "Couldn't delete the user", nil, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.ClientResponse(http.StatusOK, "User delete successfully", nil, nil))
}
