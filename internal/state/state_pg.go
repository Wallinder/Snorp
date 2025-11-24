package state

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type ArchivedMessages struct {
	ID         string `gorm:"primaryKey"`
	Type       int
	AuthorID   string
	GlobalName string
	Username   string
	Content    string
	Timestamp  time.Time
}

func (session *SessionState) CreateConnection() *gorm.DB {
	gormCfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: session.Config.Postgresql.Gorm.SingularTable,
		},
	}
	db, err := gorm.Open(postgres.Open(session.Config.Postgresql.ConnectionString), gormCfg)
	if err != nil {
		log.Fatalf("Unable to connect to postgresql: %v\n", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Unable to aquire db connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(session.Config.Postgresql.Gorm.MaxIdleConns)

	sqlDB.SetMaxOpenConns(session.Config.Postgresql.Gorm.MaxOpenConns)

	sqlDB.SetConnMaxLifetime(time.Duration(session.Config.Postgresql.Gorm.ConnMaxLifetime) * time.Second)

	err = db.AutoMigrate(&ArchivedMessages{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
