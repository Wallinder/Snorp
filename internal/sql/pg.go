package sql

import (
	"log"
	"snorp/internal/state"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Jobs struct {
	Name      string `gorm:"primaryKey"`
	Timestamp time.Time
}

type ArchivedMessages struct {
	ID         string `gorm:"primaryKey"`
	Type       int
	AuthorID   string
	GlobalName string
	Username   string
	Content    string
	Timestamp  time.Time
}

func CreateConnection(connectionString string, settings *state.DBSettings) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connectionString), settings.GormConfig)
	if err != nil {
		log.Fatalf("Unable to connect to postgresql: %v\n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Unable to aquire db connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(settings.MaxIdleConns)

	sqlDB.SetMaxOpenConns(settings.MaxOpenConns)

	sqlDB.SetConnMaxLifetime(settings.ConnMaxLifetime)

	return db
}

func InitDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&ArchivedMessages{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Jobs{})
	if err != nil {
		log.Fatal(err)
	}
}

func Insert[T any](db *gorm.DB, model *T) error {
	result := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(model)

	return result.Error
}
