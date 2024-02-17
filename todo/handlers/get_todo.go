package handlers

import (
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/bosskasidit/todo/types"
	"github.com/gin-gonic/gin"
)

func GetTodosHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		dsnap, err := client.Collection(types.TODO_COLLECTION).Doc(id).Get(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		doc := dsnap.Data()
		todo := &types.Todo{
			ID:          id,
			Title:       doc["title"].(string),
			Description: doc["description"].(string),
			CreateAt:    doc["createAt"].(time.Time),
			UpdateAt:    doc["updateAt"].(time.Time),
		}

		c.JSON(http.StatusOK, todo)
	}
}
