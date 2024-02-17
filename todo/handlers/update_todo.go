package handlers

import (
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/bosskasidit/todo/types"
	"github.com/gin-gonic/gin"
)

func UpdateTodosHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		id := c.Param("id")

		var todo types.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		todo.ID = id
		todo.UpdateAt = time.Now()
		_, err := client.
			Collection(types.TODO_COLLECTION).
			Doc(todo.ID).Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"createAt":    todo.CreateAt,
			"updateAt":    todo.UpdateAt,
		}, firestore.MergeAll)

		if err != nil {
			log.Panicf("An error has occurred: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}
func UpdateDescriptionHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		var todo types.Todo
		if err := c.BindJSON(&todo); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		if todo.Description == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "description is empty"})
		}

		if todo.ID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "id is empty"})
		}

		dsnap, err := client.Collection(types.TODO_COLLECTION).Doc(todo.ID).Get(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		todo.Title = dsnap.Data()["title"].(string)
		todo.CreateAt = dsnap.Data()["createAt"].(time.Time)
		todo.UpdateAt = time.Now()

		_, err = client.
			Collection(types.TODO_COLLECTION).
			Doc(todo.ID).Set(c, map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"createAt":    todo.CreateAt,
			"updateAt":    todo.UpdateAt,
		}, firestore.MergeAll)

		if err != nil {
			log.Panicf("An error has occurred: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}

		c.JSON(http.StatusOK, todo)
	}
}
