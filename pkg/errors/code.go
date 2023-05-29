package errors

import (
	"net/http"
)

var (
	// Common

	// ErrUnauthenticated is returned when the user is not authorized
	ErrUnauthenticated = New(http.StatusUnauthorized, 10001, "Unauthenticated.")
	// ErrBadRequest is returned when the request body is not valid
	ErrBadRequest = New(http.StatusBadRequest, 10002, "Bad request.")
	// ErrForbidden is returned when the user is not allowed to do something
	ErrForbidden = New(http.StatusForbidden, 10003, "Forbidden.")
	// ErrNotFound is returned when a resource is not found
	ErrNotFound = New(http.StatusNotFound, 10004, "Not found.")
	// ErrInternalServerError is returned when an error occurred in the application
	ErrInternalServerError = New(http.StatusInternalServerError, 10005, "Internal server error.")
	// ErrBadGateway is returned when an error occurred in the application
	ErrBadGateway = New(http.StatusBadGateway, 10006, "Bad gateway.")
	// ErrUnprocessableEntity is returned when the request body is not valid
	ErrUnprocessableEntity = New(http.StatusUnprocessableEntity, 10007, "Unprocessable entity.")

	// JWT

	// ErrJWTInvalidCredentials is returned when the user JWT credentials are invalid
	ErrJWTInvalidCredentials = New(http.StatusBadRequest, 11000, "JWT token missing or invalid.")
	// ErrJWTInvalidClaims is returned when the user JWT claims are invalid
	ErrJWTInvalidClaims = New(http.StatusBadRequest, 11001, "Failed to cast claims as jwt.MapClaims.")
	// ErrJWTRevoke is returned when the user JWT token is revoked
	ErrJWTRevoke = New(http.StatusBadRequest, 11002, "JWT token is revoked.")

	// Redis

	// ErrRedisConnection is returned when the redis connection is failed
	ErrRedisConnection = New(http.StatusInternalServerError, 12000, "Redis connection failed.")
	// ErrRedisKeyNotFound is returned when the redis key is not found
	ErrRedisKeyNotFound = New(http.StatusNotFound, 12001, "Redis key not found.")

	// Database

	// ErrUnexpectedDBError is returned when an unexpected error occurred in the database
	ErrUnexpectedDBError = New(http.StatusBadRequest, 13000, "Unexpected DB error.")

	// SendEmail

	// ErrSendEmailFromToInvalid is returned when the send email from or to is invalid
	ErrSendEmailFromToInvalid = New(
		http.StatusBadRequest,
		14000,
		"Send email must specify at least one From address and one To address.",
	)

	// Auth

	// ErrAuthLoginFailed is returned when the user login is failed
	ErrAuthLoginFailed = New(http.StatusBadRequest, 15000, "These credentials do not match our records.")
	// ErrAuthInvalidatePass is returned when the user password is invalid
	ErrAuthInvalidatePass = New(http.StatusBadRequest, 15001, "Invalid password.")
	// ErrAuthInvalidateEmail is returned when the user email is invalid
	ErrAuthInvalidateEmail = New(http.StatusBadRequest, 15002, "Invalid email.")
	// ErrAuthInvalidateConfirmPass is returned when the user confirm password is invalid
	ErrAuthInvalidateConfirmPass = New(http.StatusBadRequest, 15003, "Invalid confirm password.")
	// ErrAuthInvalidateToken is returned when the user token is invalid
	ErrAuthInvalidateToken = New(http.StatusBadRequest, 15004, "Invalid token forgot password.")
	// ErrAuthThrottleLogin is returned when the user login is failed
	ErrAuthThrottleLogin = New(http.StatusForbidden, 15005, "Too many login attempts. Please try again later.")

	// Role

	// ErrRoleExists is returned when the role already exists
	ErrRoleExists = New(http.StatusBadRequest, 16000, "Role already exists.")

	// User

	// ErrUserExistsByEmail is returned when the user already exists by email
	ErrUserExistsByEmail = New(http.StatusBadRequest, 17000, "User already exists by email.")
)
