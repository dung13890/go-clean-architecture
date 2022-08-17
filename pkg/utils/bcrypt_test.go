package utils_test

import (
	"go-app/pkg/utils"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

type ExpectedBcryptResult struct {
	pass string
	hash string
}

var expectedBcryptResults = []ExpectedBcryptResult{
	{
		pass: "verify",
		hash: "$2a$10$yC25vmGJRlYHdzcHsl8EfOIl74Jf4YqfNDgKTjAaeGSFK7qFXDAz6",
	},
	{
		pass: "verify2",
		hash: "$2a$10$4JFLFY/Y3c.YSjwDHAD.L.3wiyH2dtdhEt5Agwj.9WdQD1KS5GLzq",
	},
}

func TestGeneratePassword(t *testing.T) {
	t.Parallel()
	for testNumber, testExpected := range expectedBcryptResults {
		pass := testExpected.pass

		result, _ := utils.GeneratePassword(pass)

		if bcrypt.CompareHashAndPassword([]byte(result), []byte(pass)) != nil {
			t.Errorf("#%d (%s)\n+++ %s\n--- %s", testNumber, "TestGeneratePassword", pass, result)
		}
	}
}

func TestComparePassword(t *testing.T) {
	t.Parallel()
	for testNumber, testExpected := range expectedBcryptResults {
		pass := testExpected.pass
		hash := testExpected.hash

		if !utils.ComparePassword(pass, hash) {
			t.Errorf("#%d (%s)\n+++ %s\n--- %s", testNumber, "TestComparePassword", pass, hash)
		}
	}
}
