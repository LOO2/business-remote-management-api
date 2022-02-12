package delivery

import (
	"net/http"
	"strconv"

	"github.com/LOO2/business-remote-management-api/repository"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ShowAllUsers(c *gin.Context) {

	result, err := repository.GetAllUser()
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find User " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	result, err := repository.GetUserById(newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find User by ID: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {

	var p *repository.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.CreateUser(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create User: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, p)
}

func UpdateUser(c *gin.Context) {

	var p *repository.User

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.UpdateUser(p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot update User: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = repository.DeleteUser(repository.User{}, newid)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "cannot find User by id: " + err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
