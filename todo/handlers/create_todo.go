package handlers

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/bosskasidit/todo/types"
	"github.com/gin-gonic/gin"
)

func CreateTodoHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var todo types.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		now := time.Now()
		todo.CreateAt = now
		todo.UpdateAt = now

		ref := client.Collection(types.TODO_COLLECTION).NewDoc()
		_, err := ref.Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"completed":   false,
			"createAt":    todo.CreateAt,
			"updateAt":    todo.UpdateAt,
		})

		if err != nil {
			log.Fatalf("An error has occurred: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusCreated, todo)
	}
}
