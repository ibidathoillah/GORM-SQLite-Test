package main

import (
	"errors"

	"github.com/ibidathoillah/majoo-soal2/en"
	"github.com/ibidathoillah/majoo-soal2/models"

	"github.com/ibidathoillah/majoo-soal2/repositories"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	u := repositories.NewAreaRepository(db)
	area := models.Area{}
	err = u.InsertArea(float64(10), float64(10), models.PERSEGI, &area)
	if err != nil {
		log.Error(err)
		err = errors.New(en.ERROR_DATABASE)
	}

}
