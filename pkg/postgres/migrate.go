package postgres

import (
	"flag"
	"fmt"
	"go-app/config"
	"go-app/pkg/errors"
	"go-app/pkg/logger"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	// migrate tools
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate from db/migrates into postgres database
func Migrate(db config.Database) error {
	uri := fmt.Sprintf("%s://%s:%s@%s:%v/%s?sslmode=%s",
		db.Connection,
		db.User,
		db.Password,
		db.Host,
		strconv.Itoa(db.Port),
		db.DBName,
		db.SSLMode,
	)
	logger.Debug().Printf(
		"connect to [%v://%v:***@%v:%v/%v?sslmode=%v]",
		db.Connection,
		db.User,
		db.Host,
		db.Port,
		db.DBName,
		db.SSLMode,
	)

	m, err := migrate.New("file://db/migrations", uri)
	if err != nil {
		return errors.Wrap(err)
	}

	method := "up"

	flag.Parse()

	if flag.Arg(0) != "" {
		method = flag.Arg(0)
	}
	switch {
	case method == "up":
		up(m)
	case method == "down" && flag.Arg(1) != "":
		down(m)
	default:
		logger.Error().Fatalf("Invalid parameter: %s", method)
	}

	return nil
}

func up(m *migrate.Migrate) {
	for {
		if err := m.Steps(1); err != nil {
			if os.IsNotExist(err) {
				break
			}
			logger.Error().Fatal(err)
		}
		v, _, err := m.Version()
		if err != nil {
			logger.Error().Fatal(err)
		}
		logger.Info().Printf("Migrate up version is %v", v)
	}
}

func down(m *migrate.Migrate) {
	remain, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		logger.Error().Fatal(err)
	}
	if remain < 0 {
		logger.Error().Fatal("Down step is less than 0")
	}
	for remain > 0 {
		v, _, err := m.Version()
		if err != nil {
			logger.Error().Fatal(err)
		}
		logger.Info().Printf("Migrate down version is %v", v)
		if err := m.Steps(-1); err != nil {
			logger.Error().Fatal(err)
		}
		remain--
	}
}
