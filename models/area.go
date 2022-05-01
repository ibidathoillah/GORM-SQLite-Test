package models

type AreaType string

const (
	PERSEGI_PANJANG = "persegi_panjang"
	PERSEGI         = "persegi"
	SEGITIGA        = "segitiga"
)

type Area struct {
	ID        int64    `gorm:"column:id;primaryKey;"`
	AreaValue float64  `gorm:"column:area_value"`
	AreaType  AreaType `gorm:"column:type"`
}
