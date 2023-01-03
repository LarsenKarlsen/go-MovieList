package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInDB struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []UserInDB{
	{ID: "1", Email: "test@test", Password: "test"},
}

// add new user
func postUser(c *gin.Context) {
	var newUserInput UserInput
	var newUser UserInDB

	// read user data from request
	if err := c.BindJSON(&newUserInput); err != nil {
		return
	}
	// convert userInput to json string
	temporaryVariable, _ := json.Marshal(newUserInput)
	// read user input json to userInDB struct
	err := json.Unmarshal(temporaryVariable, &newUser)
	if err != nil {
		// Catch the exception to handle it as per your need
	}
	// add ID to user
	newUser.ID = "999"
	// save user to db
	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, newUser)
}

// add new user
func userSignIn(c *gin.Context) {

	var user UserInput

	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": err})
	}
	c.IndentedJSON(http.StatusOK, user)
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func AddUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/", getUsers)
	user.POST("/", postUser)
}
