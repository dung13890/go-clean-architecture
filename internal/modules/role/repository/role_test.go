package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"go-app/internal/domain"
	"go-app/internal/modules/role/repository"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var errRp = errors.New("Have an error")

type testcase struct {
	name string
	mock func(sqlmock.Sqlmock)
	res  interface{}
	args *domain.Role
	err  error
}

func setUp(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		Conn:                 db,
		PreferSimpleProtocol: true,
	})

	dbConnectMock, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Errorf("Failed to open gorm v2 db, got error: %v", err)
	}

	return db, dbConnectMock, mock
}

func TestFetchRole(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "roles"`,
				)
				mock.ExpectQuery(query).WithArgs().WillReturnRows(sqlmock.NewRows(nil))
			},
			res: []domain.Role{},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "roles"`,
				)
				mock.ExpectQuery(query).WithArgs().WillReturnError(errRp)
			},
			res: []domain.Role{},
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			res, err := repoMock.Fetch(context.Background())

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, res, tc.res)
		})
	}
}

func TestFindRole(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	timeNow := time.Now()
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "slug", "updated_at", "created_at", "deleted_at"}).
					AddRow(1, "name", "slug", timeNow, timeNow, nil)
				query := regexp.QuoteMeta(
					`SELECT * FROM "roles" WHERE id = $1 ORDER BY "roles"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			args: &domain.Role{
				ID: 1,
			},
			res: &domain.Role{
				ID:        1,
				Name:      "name",
				Slug:      "slug",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "roles" WHERE id = $1 ORDER BY "roles"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs(2).WillReturnError(errRp)
			},
			args: &domain.Role{
				ID: 2,
			},
			res: (*domain.Role)(nil),
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			res, err := repoMock.Find(context.Background(), int(tc.args.ID))

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, res, tc.res)
		})
	}
}

func TestStoreRole(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	timeNow := time.Now()
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				query := regexp.QuoteMeta(
					`INSERT INTO "roles" ("name","slug","created_at","updated_at","deleted_at")
								VALUES ($1,$2,$3,$4,$5) RETURNING "id"`,
				)
				mock.ExpectQuery(query).
					WithArgs("name", "name", timeNow, timeNow, nil).
					WillReturnRows(rows)
			},
			args: &domain.Role{
				Name:      "name",
				Slug:      "name",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`INSERT INTO "roles" ("name","slug","created_at","updated_at","deleted_at")
								VALUES ($1,$2,$3,$4,$5) RETURNING "id"`,
				)
				mock.ExpectQuery(query).
					WithArgs("name2", "name2", timeNow, timeNow, nil).
					WillReturnError(errRp)
			},
			args: &domain.Role{
				Name:      "name2",
				Slug:      "name2",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			err := repoMock.Store(context.Background(), tc.args)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}
