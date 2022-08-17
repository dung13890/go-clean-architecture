package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"go-app/internal/domain"
	"go-app/internal/modules/user/repository"
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
	args interface{}
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

func TestFetchUser(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "users"`,
				)
				mock.ExpectQuery(query).WithArgs().WillReturnRows(sqlmock.NewRows(nil))
			},
			res: []domain.User{},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "users"`,
				)
				mock.ExpectQuery(query).WithArgs().WillReturnError(errRp)
			},
			res: []domain.User{},
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

func TestFindUser(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	timeNow := time.Now()
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "role_id", "created_at", "updated_at", "deleted_at"}).
					AddRow(1, "name", "email", 1, timeNow, timeNow, nil)
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)
			},
			args: &domain.User{
				ID: 1,
			},
			res: &domain.User{
				ID:        1,
				Name:      "name",
				Email:     "email",
				RoleID:    1,
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE id = $1 ORDER BY "users"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs(2).WillReturnError(errRp)
			},
			args: &domain.User{
				ID: 2,
			},
			res: (*domain.User)(nil),
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			res, err := repoMock.Find(context.Background(), int(tc.args.(*domain.User).ID))

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, res, tc.res)
		})
	}
}

func TestStoreUser(t *testing.T) {
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
					`INSERT INTO "users" ("name","email","role_id","password","created_at","updated_at","deleted_at")
								VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`,
				)
				mock.ExpectQuery(query).
					WithArgs("name", "email", 1, sqlmock.AnyArg(), timeNow, timeNow, nil).
					WillReturnRows(rows)
			},
			args: &domain.User{
				Name:      "name",
				Email:     "email",
				RoleID:    1,
				Password:  "password",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`INSERT INTO "users" ("name","email","role_id","password","created_at","updated_at","deleted_at")
								VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING "id"`,
				)
				mock.ExpectQuery(query).
					WithArgs("name2", "email2", 2, sqlmock.AnyArg(), timeNow, timeNow, nil).
					WillReturnError(errRp)
			},
			args: &domain.User{
				Name:      "name2",
				Email:     "email2",
				RoleID:    2,
				Password:  "password",
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
			args, _ := tc.args.(*domain.User)
			err := repoMock.Store(context.Background(), args)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestSearchUser(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	timeNow := time.Now()
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id",
					"name",
					"email",
					"role_id",
					"password",
					"created_at",
					"updated_at",
					"deleted_at",
				}).
					AddRow(1, "name", "email", 1, "password", timeNow, timeNow, nil)
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1`,
				)
				mock.ExpectQuery(query).WithArgs("email").WillReturnRows(rows)
			},
			args: domain.UserQueryParam{
				Email: "email",
			},
			res: []domain.User{
				{
					ID:        1,
					Name:      "name",
					Email:     "email",
					RoleID:    1,
					Password:  "password",
					CreatedAt: timeNow,
					UpdatedAt: timeNow,
				},
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1`,
				)
				mock.ExpectQuery(query).WithArgs("email2").WillReturnError(errRp)
			},
			args: domain.UserQueryParam{
				Email: "email2",
			},
			res: []domain.User(nil),
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			res, err := repoMock.Search(context.Background(), tc.args.(domain.UserQueryParam))

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, res, tc.res)
		})
	}
}

func TestFindByQuery(t *testing.T) {
	t.Parallel()
	db, dbConnectMock, mock := setUp(t)
	timeNow := time.Now()
	defer db.Close()

	repoMock := repository.NewRepository(dbConnectMock)

	tcs := []testcase{
		{
			name: "OK",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id",
					"name",
					"email",
					"role_id",
					"password",
					"created_at",
					"updated_at",
					"deleted_at",
				}).
					AddRow(1, "name", "email", 1, "password", timeNow, timeNow, nil)
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs("email").WillReturnRows(rows)
			},
			args: domain.UserQueryParam{
				Email: "email",
			},
			res: &domain.User{
				ID:        1,
				Name:      "name",
				Email:     "email",
				RoleID:    1,
				Password:  "password",
				CreatedAt: timeNow,
				UpdatedAt: timeNow,
			},
			err: nil,
		},
		{
			name: "NG",
			mock: func(mock sqlmock.Sqlmock) {
				query := regexp.QuoteMeta(
					`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."id" LIMIT 1`,
				)
				mock.ExpectQuery(query).WithArgs("email2").WillReturnError(errRp)
			},
			args: domain.UserQueryParam{
				Email: "email2",
			},
			res: (*domain.User)(nil),
			err: errRp,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.mock(mock)
			res, err := repoMock.FindByQuery(context.Background(), tc.args.(domain.UserQueryParam))

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, res, tc.res)
		})
	}
}
