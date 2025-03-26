package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []user{}

func main() {
	router := gin.Default()
	router.GET("/register", getRegisters)
	router.GET("/register/:id", getRegisterByID)
	router.POST("/register", postRegister)

	router.Run("0.0.0.0:8080")
}

func Register(user user) {
	users = append(users, user)
}

func generatePass(user user) user {
	user.Password = "123456"
	return user
}

func getRegisters(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getRegisterByID(c *gin.Context) {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func postRegister(c *gin.Context) {
	var newUser user
	res := c.BindJSON(&newUser)
	if res != nil {
		return
	}
	for _, u := range users {
		if u.ID == newUser.ID {
			c.IndentedJSON(http.StatusConflict, gin.H{"message": "id ja cadastrado"})
			return
		}
	}
	userWithPass := generatePass(newUser)
	Register(userWithPass)
	c.IndentedJSON(http.StatusCreated, newUser)

}
