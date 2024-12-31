package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Put(c *gin.Context) {
	article := Article{}

	c.JSON(http.StatusOK, article)
}
