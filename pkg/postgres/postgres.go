package postgres

import (
	"fmt"
	"go-app/config"
	"go-app/pkg/errors"
	"strconv"

	// for postgresql driver.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewGormDB setup Gorm
func NewGormDB(db config.Database) (*gorm.DB, error) {
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Host,
		strconv.Itoa(db.Port),
		db.User,
		db.Password,
		db.DBName,
		db.SSLMode,
	)

	dbConnect, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return dbConnect, nil
}
