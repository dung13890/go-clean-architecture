package usecase_test

import (
	"context"
	"errors"
	"go-app/internal/domain"
	mockDomain "go-app/internal/domain/mock"
	"go-app/internal/modules/auth/usecase"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

var errUserUc = errors.New("Have an error")

type userTestcase struct {
	name string
	mock func(*mockDomain.MockUserRepository)
	res  interface{}
	args interface{}
	err  error
}

func TestFetchUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(repo)

	tests := []userTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Fetch(context.Background()).Return(nil, nil)
			},
			res: []domain.User(nil),
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)

			res, err := uc.Fetch(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestFindUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(repo)

	tests := []userTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Find(context.Background(), 1).Return(&domain.User{}, nil)
			},
			res:  &domain.User{},
			args: 1,
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Find(context.Background(), 2).Return(nil, errUserUc)
			},
			res:  (*domain.User)(nil),
			args: 2,
			err:  errUserUc,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)

			res, err := uc.Find(context.Background(), tc.args.(int))

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}

func TestStoreUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []userTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), &domain.User{}).Return(nil)
			},
			args: &domain.User{},
			res:  nil,
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), &domain.User{Name: "name"}).Return(errUserUc)
			},
			args: &domain.User{Name: "name"},
			res:  nil,
			err:  errUserUc,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			repo := mockDomain.NewMockUserRepository(ctrl)
			uc := usecase.NewUserUsecase(repo)

			tc.mock(repo)
			args, _ := tc.args.(*domain.User)

			res := uc.Store(context.Background(), args)
			require.ErrorIs(t, res, tc.err)
		})
	}
}

func TestFindByQueryUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockUserRepository(ctrl)
	uc := usecase.NewUserUsecase(repo)

	tests := []userTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.User{Email: "email"}).
					Return(&domain.User{}, nil)
			},
			res:  &domain.User{},
			args: domain.User{Email: "email"},
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.User{Email: "email2"}).
					Return(nil, errUserUc)
			},
			res:  (*domain.User)(nil),
			args: domain.User{Email: "email2"},
			err:  errUserUc,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)

			res, err := uc.FindByQuery(context.Background(), tc.args.(domain.User))

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}
}
