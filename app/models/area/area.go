package area

import (
	"gigphil/app/models/base"
	"time"
)

type Area struct {
	base.Model
	ID        int64 `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func All() []Area {
	var areas []Area
	db := Area{}.GetDB()
	db.Find(&areas)
	return areas
}

func AllNames() []string {
	var names []string
	db := Area{}.GetDB()
	db.Model(&Area{}).Pluck("name", &names)
	return names
}
