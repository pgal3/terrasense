package pg_adapter

import (
	"fmt"

	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBClient(config PostgreSQLConfig) *gorm.DB {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Url, config.User, config.Pwd, config.DBName, config.Port)
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Setup Tables
	db.AutoMigrate(&pg_models.Measurement{}, &pg_models.Pairing{})
	//Setup conn pool
	// pg , _:= db.DB()
	// pg.SetMaxIdleConns(5)
	// pg.SetMaxOpenConns(10)
	return db
}