package controllers

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/naewatcharapong/BeerAPItest/logger"
	"github.com/naewatcharapong/BeerAPItest/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BeerRepo struct {
	Db     *gorm.DB
	Logger *logger.MongoLogger
}

func New(db *gorm.DB, mongologger *logger.MongoLogger) *BeerRepo {
	db.AutoMigrate(&models.Beer{})
	return &BeerRepo{Db: db, Logger: mongologger}
}

//create Beer
func (repository *BeerRepo) InsertBeer(c *gin.Context) {
	var BeerModel models.BeerModel
	var Beer models.Beer
	c.BindJSON(&BeerModel)
	sDec, _ := b64.StdEncoding.DecodeString(BeerModel.Images)
	Beer.Images = sDec
	err := models.InsertBeer(repository.Db, &Beer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	repository.Logger.Log(&Beer, "INSERT")
	c.JSON(http.StatusOK, Beer)
}

//get Beers
func (repository *BeerRepo) GetBeers(c *gin.Context) {
	var Beer []models.Beer
	name := c.Query("name")
	err := models.GetBeers(repository.Db, &Beer, name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Beer)
}

//get Beer by id
func (repository *BeerRepo) GetBeer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Beer models.Beer
	err := models.GetBeer(repository.Db, &Beer, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Beer)
}

// update Beer
func (repository *BeerRepo) UpdateBeer(c *gin.Context) {
	var BeerModel models.BeerModel
	var Beer models.Beer
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.GetBeer(repository.Db, &Beer, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&BeerModel)

	sDec, _ := b64.StdEncoding.DecodeString(BeerModel.Images)
	fmt.Println(sDec)
	Beer.Images = sDec
	err = models.UpdateBeer(repository.Db, &Beer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	repository.Logger.Log(&Beer, "UPDATE")
	c.JSON(http.StatusOK, BeerModel)
}

// delete Beer
func (repository *BeerRepo) DeleteBeer(c *gin.Context) {
	var Beer models.Beer
	id, _ := strconv.Atoi(c.Param("id"))
	err := models.DeleteBeer(repository.Db, &Beer, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	repository.Logger.Log(&Beer, "DELETE")
	c.JSON(http.StatusOK, gin.H{"message": "Beer deleted successfully"})
}
