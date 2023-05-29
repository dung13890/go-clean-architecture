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
	errAuthUc             = errors.New("Have an error")
	errAuthInvalidatePass = errors.New("invalidate Password")
)

type authTestcase struct {
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

	jwtSvc := mockDomain.NewMockJWTService(ctrl)
	thSvc := mockDomain.NewMockThrottleService(ctrl)
	repo := mockDomain.NewMockUserRepository(ctrl)
	pwRepo := mockDomain.NewMockPasswordResetRepository(ctrl)
	uc := usecase.NewAuthUsecase(jwtSvc, thSvc, repo, pwRepo)

	tests := []authTestcase{
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
				repo.EXPECT().Store(context.Background(), &domain.User{Name: "1"}).Return(errAuthUc)
			},
			args: &domain.User{Name: "1"},
			res:  (*domain.User)(nil),
			err:  errAuthUc,
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

	jwtSvc := mockDomain.NewMockJWTService(ctrl)
	thSvc := mockDomain.NewMockThrottleService(ctrl)
	repo := mockDomain.NewMockUserRepository(ctrl)
	pwRepo := mockDomain.NewMockPasswordResetRepository(ctrl)
	uc := usecase.NewAuthUsecase(jwtSvc, thSvc, repo, pwRepo)

	tests := []authTestcase{
		{
			name: "OK",
			mock: func(repo *mockDomain.MockUserRepository) {
				passHash, _ := utils.GeneratePassword("")
				jwtSvc.EXPECT().GenerateToken(context.Background(), gomock.Any()).Return("hash_token1", int64(0), nil)
				thSvc.EXPECT().Blocked(context.Background(), "email@email.com", "ip").Return(false, nil)
				thSvc.EXPECT().Clear(context.Background(), "email@email.com", "ip").Return(nil)
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.User{Email: "email@email.com"}).
					Return(&domain.User{Email: "email@email.com", Password: passHash}, nil)
			},
			args: &domain.User{Email: "email@email.com"},
			res:  "hash_token1",
			err:  nil,
		},
		{
			name: "FindByQuery NG",
			mock: func(repo *mockDomain.MockUserRepository) {
				thSvc.EXPECT().Blocked(context.Background(), "email1@email.com", "ip").Return(false, nil)
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.User{Email: "email1@email.com"}).
					Return(nil, errAuthUc)
			},
			args: &domain.User{Email: "email1@email.com"},
			res:  "",
			err:  errAuthUc,
		},
		{
			name: "ComparePassword NG",
			mock: func(repo *mockDomain.MockUserRepository) {
				passHash, _ := utils.GeneratePassword("wrong!")
				thSvc.EXPECT().Blocked(context.Background(), "email2@email.com", "ip").Return(false, nil)
				thSvc.EXPECT().Incr(context.Background(), "email2@email.com", "ip").Return(nil)
				repo.
					EXPECT().
					FindByQuery(context.Background(), domain.User{Email: "email2@email.com"}).
					Return(&domain.User{Password: passHash}, nil)
			},
			args: &domain.User{Email: "email2@email.com"},
			res:  "",
			err:  errAuthInvalidatePass,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			tc.mock(repo)
			args, _ := tc.args.(*domain.User)
			token, exp, err := uc.Login(context.Background(), args, "ip")
			assert.Equal(t, token, tc.res)
			assert.Equal(t, exp, int64(0))
			if tc.err != nil {
				assert.Errorf(t, err, tc.err.Error())
			}
		})
	}
}
