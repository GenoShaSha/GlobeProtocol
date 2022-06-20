package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

var users = []user{
	{Id: 0, FirstName: "Shanessa", LastName: "Kostaman", Username: "gnvshanessa", Email: "shanessa.m7493@gmail.com", Password: "shanessa12345", PhoneNumber: "0657854265"},
	{Id: 1, FirstName: "Hristo", LastName: "Hristov", Username: "HristoHristov", Email: "hristo2000@gmail.com", Password: "hristo12345", PhoneNumber: "0654382195"},
	{Id: 2, FirstName: "Dobri", LastName: "Trifonov", Username: "DobriTrifonov", Email: "dobriTrifonov@gmail.com", Password: "dobri12345", PhoneNumber: "06574328443"},
}

func main() {
	router := gin.Default()
	router.GET("/User", getUsers)
	router.GET("/User/:email", getSelectedUser)
	router.PATCH("/User/:username", updateUser)
	router.DELETE("/User/:username", removeSelectedUser)
	router.POST("/SignUp", addUser)
	router.Run("localhost:7777")
}

func addUser(context *gin.Context) {
	var addNewUser user

	if err := context.BindJSON(&addNewUser); err != nil {
		return

	}
	users = append(users, addNewUser)
	context.IndentedJSON(http.StatusCreated, addNewUser)
}

func getUsers(context *gin.Context) {
	context.JSON(http.StatusOK, users)
}

func getSelectedUser(context *gin.Context) {
	email := context.Param("email")

	user, err := getUserByEmail(email)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the user!"})
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserByEmail(email string) (*user, error) {
	for i, usr := range users {
		if usr.Email == email {
			return &users[i], nil
		}
	}
	return nil, errors.New("Couldn't find any user with the given email!")
}

func updateUser(context *gin.Context) {

	Username := context.Param("username")

	tempUser, err := getUserByUsername(Username)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the user!"})
	}

	newUser := user{}
	context.BindJSON(&newUser)

	tempUser.FirstName = newUser.FirstName
	tempUser.LastName = newUser.LastName
	tempUser.Email = newUser.Email
	tempUser.Password = newUser.Password
	tempUser.PhoneNumber = newUser.PhoneNumber
	context.IndentedJSON(http.StatusOK, newUser)
}

func getUserByUsername(username string) (*user, error) {
	for i, usr := range users {
		if usr.Username == username {
			return &users[i], nil
		}
	}
	return nil, errors.New("Couldn't find any user with this username!")
}

func removeSelectedUser(context *gin.Context) {
	username := context.Param("username")

	userr, err := getUserByUsername(username)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Couldn't find the user!"})
	}
	users = remove(users, userr.Id)
	context.IndentedJSON(http.StatusOK, userr)
}

func remove(u []user, index int) []user {
	return append(u[:index], u[index+1:]...)
}
