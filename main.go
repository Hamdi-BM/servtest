package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]*User{}

func main() {
	loadUsersFromFile("users.json")
	router := echo.New()
	router.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "welcome to the server")
	})

	router.GET("/users/:id", getUser)
	router.POST("/users", saveUser)
	router.DELETE("/users/:id", deleteUser)

	router.Logger.Fatal(router.Start(":8088"))
}
func getUser(c echo.Context) error {
	id := c.Param("id")
	userID := 0
	fmt.Sscanf(id, "%d", &userID)

	user, ok := users[userID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)

}

func loadUsersFromFile(filename string) {
	// Reading JSON data mil file aka users.json
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON data fil map hathi: var users = map[int]*User{}
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}
func saveUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.ID = len(users) + 1
	users[user.ID] = user
	return c.JSON(http.StatusCreated, user)

}
func deleteUser(c echo.Context) error {
	id := c.Param("id")
	userID := 0
	fmt.Sscanf(id, "%d", &userID)

	_, ok := users[userID]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	delete(users, userID)
	return c.NoContent(http.StatusNoContent)
}
