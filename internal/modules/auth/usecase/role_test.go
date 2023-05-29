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

var errRoleUc = errors.New("Have an error")

type roleTestcase struct {
	name string
	mock func(*mockDomain.MockRoleRepository)
	res  interface{}
	args interface{}
	err  error
}

func TestFetchRole(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockRoleRepository(ctrl)
	uc := usecase.NewRoleUsecase(repo)

	tests := []roleTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockRoleRepository) {
				repo.EXPECT().Fetch(context.Background()).Return(nil, nil)
			},
			res: []domain.Role(nil),
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

func TestFindRole(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockRoleRepository(ctrl)
	uc := usecase.NewRoleUsecase(repo)

	tests := []roleTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockRoleRepository) {
				repo.EXPECT().Find(context.Background(), 1).Return(&domain.Role{}, nil)
			},
			res:  &domain.Role{},
			args: 1,
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockRoleRepository) {
				repo.EXPECT().Find(context.Background(), 2).Return(nil, errRoleUc)
			},
			res:  (*domain.Role)(nil),
			args: 2,
			err:  errRoleUc,
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

func TestStoreRole(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []roleTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockRoleRepository) {
				repo.EXPECT().Store(context.Background(), &domain.Role{}).Return(nil)
			},
			args: &domain.Role{},
			res:  nil,
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockRoleRepository) {
				repo.EXPECT().Store(context.Background(), &domain.Role{}).Return(errRoleUc)
			},
			res: nil,
			err: errRoleUc,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			repo := mockDomain.NewMockRoleRepository(ctrl)
			uc := usecase.NewRoleUsecase(repo)

			tc.mock(repo)

			res := uc.Store(context.Background(), &domain.Role{})
			require.ErrorIs(t, res, tc.err)
		})
	}
}
