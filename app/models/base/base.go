package base

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Model struct {
	gorm.Model
}

func (model Model) GetDB() *gorm.DB {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger = db.Logger.LogMode(logger.Info)
	if err != nil {
		panic("failed to connect database")
	}
	// egormではdeleted_atが必須になってしまい、sql流す時に以下クエリが追加されてしまうためUnscopeしている
	// WHERE "deleted_at" IS NULL
	return db.Unscoped()
}
