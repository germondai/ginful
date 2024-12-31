package main

import (
	"ginful/article"
	"ginful/user"
	"ginful/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", index)

	router.GET("/article", article.GetAll)
	router.GET("/article/:id", article.Get)

	router.GET("/user", user.GetAll)
	router.GET("/user/:id", user.Get)

	err := router.Run()
	utils.HandleErr(err)
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}
