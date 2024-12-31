package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(ids ...string) []User {
	if len(ids) == 0 {
		return Users
	}

	var result []User
	idSet := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	for _, User := range Users {
		if _, exists := idSet[User.ID]; exists {
			result = append(result, User)
		}
	}

	return result
}

func GetAll(c *gin.Context) {
	Users := get()
	c.JSON(http.StatusOK, Users)
}

func Get(c *gin.Context) {
	id := c.Param("id")
	Users := get(id)

	if len(Users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, Users[0])
}
