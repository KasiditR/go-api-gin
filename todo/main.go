package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bosskasidit/todo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := createClient()

	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	r := gin.Default()

	r.GET("/api/health", handlers.HealthCheckHandler())
	r.POST("/api/create-todo", handlers.CreateTodoHandler(client))
	r.GET("/api/get-all-todos", handlers.ListTodosHandler(client))
	r.GET("/api/get-todo/:id", handlers.GetTodosHandler(client))

	r.PATCH("/api/update-todo/:id", handlers.UpdateTodosHandler(client))
	r.DELETE("/api/delete-todo/:id", handlers.DeleteTodosHandler(client))

	r.Run(":8080")
}

func createClient() (*firestore.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "go-deploy-414616"}

	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalf("error init app: %v", err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalf("error when creating client: %v", err)
		return nil, err
	}
	return client, nil
}
