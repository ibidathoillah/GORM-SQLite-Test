package repositories

import (
	"github.com/ibidathoillah/majoo-soal2/models"
	"gorm.io/gorm"
)

type areaRepository struct {
	db *gorm.DB
}

type AreaRepositoryUseCase interface {
	InsertArea(param1 float64, param2 float64, param3 string, ar *models.Area) (err error)
}

func (r *areaRepository) InsertArea(param1 float64, param2 float64, param3 string, ar *models.Area) (err error) {
	var area float64
	switch param3 {
	case models.PERSEGI_PANJANG:
		area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = models.PERSEGI_PANJANG
		if err := r.db.Create(&ar).Error; err != nil {
			return err
		}
	case models.PERSEGI:
		var area = param1 * param2
		ar.AreaValue = area
		ar.AreaType = models.PERSEGI
		if err := r.db.Create(&ar).Error; err != nil {
			return err
		}
	case models.SEGITIGA:
		area = float64(0.5) * param1 * param2
		ar.AreaValue = area
		ar.AreaType = models.SEGITIGA
		if err := r.db.Create(&ar).Error; err != nil {
			return err
		}
	default:
		ar.AreaValue = 0
		ar.AreaType = models.AreaType("")
		if err := r.db.Create(&ar).Error; err != nil {
			return err
		}
	}

	return err
}

func NewAreaRepository(
	db *gorm.DB,
) AreaRepositoryUseCase {
	return &areaRepository{
		db: db,
	}
}
