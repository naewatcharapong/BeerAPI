package models

import (
	"gorm.io/gorm"
)

type Beer struct {
	ID uint
	Name   string `json:"name"`
	Type   string `json:"type"`
	Detail string `json:"detail"`
	Images []byte `json:"images"`
	gorm.Model
}

//create a Beer
func InsertBeer(db *gorm.DB, Beer *Beer) (err error) {
	err = db.Create(Beer).Error
	if err != nil {
		return err
	}
	return nil
}

//get Beers
func GetBeers(db *gorm.DB, Beer *[]Beer,name string) (err error) {
	if name != "" {
		err = db.Where("name = ?", name).Find(Beer).Error
			if err != nil {
			return err
		}
	} else {
	err = db.Find(Beer).Error
	if err != nil {
		return err
	}
}
	return nil
}

//get Beer by id
func GetBeer(db *gorm.DB, Beer *Beer, id int) (err error) {

	err = db.Where("id = ?", id).First(Beer).Error
	if err != nil {
		return err
	}
	return nil
}

//update Beer
func UpdateBeer(db *gorm.DB, Beer *Beer) (err error) {
	db.Save(Beer)
	return nil
}

//delete Beer
func DeleteBeer(db *gorm.DB, Beer *Beer, id int) (err error) {
	db.Where("id = ?", id).Delete(Beer)
	return nil
}
