package http

import (
	"go-app/internal/domain"
	"go-app/pkg/logger"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	jwt "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// JwtKey is secret key fow singed
var JwtKey = []byte("go-clean-architecture")

// Claims is struct claims for jwt
type Claims struct {
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	RoleID int    `json:"role_id"`
	jwt.StandardClaims
}

// AuthHandler represent the httphandler
type AuthHandler struct {
	Usecase domain.UserUsecase
}

// GenerateToken returns token string
func GenerateToken(user *domain.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)

	claims := &Claims{
		Name:   user.Name,
		ID:     user.ID,
		RoleID: user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(JwtKey)
	if err != nil {
		logger.Error().Println("error when new claims: ", err)

		return "", err
	}

	return tokenString, nil
}

// GeneratePassword returns hashed password
func GeneratePassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

	return string(bytes), err
}

// ComparePassword used to compare password with hashed password
func ComparePassword(pass string, hashPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(pass)) == nil
}

// NewAuthHandler will initialize the Auth endpoint
func NewAuthHandler(g *echo.Group, uc domain.UserUsecase) {
	handler := &AuthHandler{
		Usecase: uc,
	}

	g.POST("/login", handler.Login)
	g.POST("/register", handler.Register)
}

// Login for user
func (hl *AuthHandler) Login(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	query := domain.QueryParam{Email: user.Email}
	UserDb, err := hl.Usecase.FindByQuery(ctx, query)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: "Invalidate Email"})
	}

	if !ComparePassword(user.Password, UserDb.Password) {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: "Invalidate Password"})
	}

	tokenStr, err := GenerateToken(UserDb)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, UserLoginResponse{
		UserID:      UserDb.ID,
		Email:       UserDb.Email,
		RoleID:      UserDb.RoleID,
		AccessToken: tokenStr,
	})
}

// Register for user
func (hl *AuthHandler) Register(c echo.Context) error {
	user := &domain.User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, &errorResponse{Message: err.Error()})
	}

	ctx := c.Request().Context()
	passHash, err := GeneratePassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	user.Password = passHash
	err = hl.Usecase.Store(ctx, user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &errorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}
