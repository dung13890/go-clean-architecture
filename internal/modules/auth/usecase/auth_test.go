package usecase_test

import (
	"context"
	"errors"
	"go-app/internal/domain"
	mockDomain "go-app/internal/domain/mock"
	"go-app/internal/modules/auth/usecase"
	"go-app/pkg/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	errUc             = errors.New("Have an error")
	errInvalidatePass = errors.New("invalidate Password")
)

type test struct {
	name string
	mock func(*mockDomain.MockUserRepository)
	res  interface{}
	args interface{}
	err  error
}

func TestRegisterAuth(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockUserRepository(ctrl)
	uc := usecase.NewUsecase(repo)

	tests := []test{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), &domain.User{}).Return(nil)
			},
			args: &domain.User{},
			res:  &domain.User{},
			err:  nil,
		},
		{
			name: "Not Good",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.EXPECT().Store(context.Background(), &domain.User{Name: "1"}).Return(errUc)
			},
			args: &domain.User{Name: "1"},
			res:  (*domain.User)(nil),
			err:  errUc,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)

			res, err := uc.Register(context.Background(), tc.args.(*domain.User))
			assert.Equal(t, res, tc.res)
			assert.ErrorIs(t, err, tc.err)
		})
	}
}

func TestLoginAuth(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mockDomain.NewMockUserRepository(ctrl)
	uc := usecase.NewUsecase(repo)

	tests := []test{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				passHash, _ := utils.GeneratePassword("")
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.UserQueryParam{Email: "email@email.com"}).
					Return(&domain.User{Password: passHash}, nil)
			},
			args: &domain.User{Email: "email@email.com"},
			res:  &domain.Claims{},
			err:  nil,
		},
		{
			name: "FindByQuery NG",
			mock: func(repo *mockDomain.MockUserRepository) {
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.UserQueryParam{Email: "email1@email.com"}).
					Return(nil, errUc)
			},
			args: &domain.User{Email: "email1@email.com"},
			res:  (*domain.Claims)(nil),
			err:  errUc,
		},
		{
			name: "ComparePassword NG",
			mock: func(repo *mockDomain.MockUserRepository) {
				passHash, _ := utils.GeneratePassword("wrong!")
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.UserQueryParam{Email: "email2@email.com"}).
					Return(&domain.User{Password: passHash}, nil)
			},
			args: &domain.User{Email: "email2@email.com"},
			res:  (*domain.Claims)(nil),
			err:  errInvalidatePass,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)
			args, _ := tc.args.(*domain.User)
			res, _, err := uc.Login(context.Background(), args)
			if res != nil {
				res.StandardClaims.ExpiresAt = 0
			}
			assert.Equal(t, res, tc.res)
			if tc.err != nil {
				assert.Errorf(t, err, tc.err.Error())
			}
		})
	}
}
