package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Patch(c *gin.Context) {
	article := Article{}

	c.JSON(http.StatusOK, article)
}
