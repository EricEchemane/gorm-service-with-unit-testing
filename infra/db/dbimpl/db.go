package dbimpl

import (
	"fmt"
	"gopher/infra/db"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	gorm.DB
}

func New(dst ...interface{}) db.IDB {
	dsn := "host=localhost user=postgres password=19126222 dbname=gopher port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(dst...)
	return &database{
		*db,
	}

}

func (db *database) RawScan(sql string, dest interface{}, values ...interface{}) error {
	err := db.Raw(sql, values...).Scan(dest).Error
	fmt.Println("real db rawscan called")
	if err != nil {
		return err
	}
	return nil
}
