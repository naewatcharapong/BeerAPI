package logger

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/naewatcharapong/BeerAPItest/models"
)

type MongoLogger struct {
	MongoCollection *mongo.Collection
}

func New(mongoCollection *mongo.Collection) *MongoLogger {
	return &MongoLogger{MongoCollection: mongoCollection}
}

type MongoDocument struct {
	ID       uint
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Detail   string    `json:"detail"`
	ActionAt time.Time `json:"action_at"`
	Action   string    `json:"action"`
}

func (Logger *MongoLogger) Log(beer *models.Beer, action string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	time_action, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	newMongoDocument := MongoDocument{
		ID:       beer.ID,
		Name:     beer.Name,
		Type:     beer.Type,
		Detail:   beer.Detail,
		ActionAt: time_action,
		Action:   action,
	}
	result, err := Logger.MongoCollection.InsertOne(ctx, newMongoDocument)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
