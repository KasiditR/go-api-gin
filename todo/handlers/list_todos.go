package handlers

import (
	"net/http"
	"time"

	"github.com/bosskasidit/todo/types"
	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func ListTodosHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var result []types.Todo
		iter := client.Collection(types.TODO_COLLECTION).Documents(c)

		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}

			result = append(result, types.Todo{
				ID:          doc.Ref.ID,
				Title:       doc.Data()["title"].(string),
				Description: doc.Data()["description"].(string),
				CreateAt:    doc.Data()["createAt"].(time.Time),
				UpdateAt:    doc.Data()["updateAt"].(time.Time),
			})
		}

		c.JSON(http.StatusOK, result)
	}
}
