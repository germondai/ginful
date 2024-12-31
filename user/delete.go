package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	user := User{}

	c.JSON(http.StatusOK, user)
}
