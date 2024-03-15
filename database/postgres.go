package database

import (
	"fmt"

	"github.com/nattapat-w/chatapp/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.DBName,
		cfg.Db.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &postgresDB{Db: db}
}

// implement database GetDb
func (p *postgresDB) GetDb() *gorm.DB {
	return p.Db
}
