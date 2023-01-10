package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/naewatcharapong/BeerAPItest/config"
	"github.com/naewatcharapong/BeerAPItest/controllers"
	"github.com/naewatcharapong/BeerAPItest/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	r := setupRouter()

	r.Run() //deafult port 8080s
}

var LoggerCollection *mongo.Collection = config.GetCollection(config.DB, "Loggers")

func setupRouter() *gin.Engine {
	r := gin.Default()
	db := config.InitDb()
	config.ConnectDB()
	var LoggerCollection *mongo.Collection = config.GetCollection(config.DB, "Loggers")
	mongologger := logger.New(LoggerCollection)
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := controllers.New(db, mongologger)
	r.POST("/beer", userRepo.InsertBeer)
	r.GET("/beer", userRepo.GetBeers)
	r.GET("/beer/:id", userRepo.GetBeer)
	r.PUT("/beer/:id", userRepo.UpdateBeer)
	r.DELETE("/beer/:id", userRepo.DeleteBeer)
	return r
}
