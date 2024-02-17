package handlers

import (
	"cloud.google.com/go/firestore"
	"github.com/bosskasidit/todo/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteTodosHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		_, err := client.Collection(types.TODO_COLLECTION).Doc(id).Delete(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusOK, "")
	}
}
