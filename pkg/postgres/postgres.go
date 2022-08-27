package postgres

import (
	"fmt"
	"go-app/config"
	"go-app/pkg/errors"
	"strconv"

	// for postgresql driver.
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	logLevel := logger.Silent
	if db.Debug {
		// I use an env variable LOG_SQL to set logSql to either true or false.
		logLevel = logger.Info
	}

	dbConnect, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return dbConnect, nil
}
