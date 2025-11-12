package state

import (
	"context"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Jobs struct {
	Name      string `gorm:"primaryKey"`
	Timestamp time.Time
}

type Cluster struct {
	Instance  int `gorm:"primaryKey"`
	Shards    int
	Heartbeat time.Time
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

	return db
}

func (session *SessionState) InitDatabase() {
	err := session.DB.AutoMigrate(&Cluster{})
	if err != nil {
		log.Fatal(err)
	}
	if session.DB.Migrator().HasTable(&Cluster{}) {
		delete := session.DB.Where("heartbeat < now() - INTERVAL '60 minutes'").Delete(&Cluster{})
		if delete.Error != nil {
			log.Fatal(err)
		}
		log.Printf("Removed %d, stale instances from database\n", delete.RowsAffected)
	}

	err = session.DB.AutoMigrate(&ArchivedMessages{})
	if err != nil {
		log.Fatal(err)
	}

	err = session.DB.AutoMigrate(&Jobs{})
	if err != nil {
		log.Fatal(err)
	}
}

func (session *SessionState) ClusterOperator(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-ticker.C:
			var cluster []Cluster

			find := session.DB.Where("heartbeat < now() - INTERVAL '1 minutes'").Find(&cluster)
			if find.Error != nil {
				log.Fatal(find.Error)
			}

			delete := session.DB.Delete(&cluster)
			if delete.Error != nil {
				log.Fatal(delete.Error)
			}
		}
	}
}
