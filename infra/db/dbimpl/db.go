package dbimpl

import (
	"fmt"
	"gopher/infra/db"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	gorm.DB
}

func New(dst ...interface{}) db.IDB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USERNAME"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("⭕ Failed to connect database")
	}

	log.Println("✅ Connected to database")
	db.AutoMigrate(dst...)
	return &database{
		*db,
	}

}

func (db *database) RawScan(sql string, dest interface{}, values ...interface{}) error {
	err := db.Raw(sql, values...).Scan(dest).Error
	if err != nil {
		return err
	}
	return nil
}
