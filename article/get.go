package article

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func get(ids ...string) []Article {
	if len(ids) == 0 {
		return Articles
	}

	var result []Article
	idSet := make(map[string]struct{}, len(ids))
	for _, id := range ids {
		idSet[id] = struct{}{}
	}

	for _, article := range Articles {
		if _, exists := idSet[article.ID]; exists {
			result = append(result, article)
		}
	}

	return result
}

func GetAll(c *gin.Context) {
	articles := get()
	c.JSON(http.StatusOK, articles)
}

func Get(c *gin.Context) {
	id := c.Param("id")
	articles := get(id)

	if len(articles) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Article not found",
		})
		return
	}

	c.JSON(http.StatusOK, articles[0])
}
