package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var password = "abcd12345"
var hasedPassword string
var errors error

func TestHashPassword(t *testing.T) {
	hasedPassword, errors = HashPassword(password)
	assert.NoError(t, errors, "error occured while hasing password")
	assert.NotEqual(t, password, hasedPassword, "Password hashed successfully")
}

func TestVerifyPassword(t *testing.T) {
	t.Run("verify success", func(t *testing.T) {
		var err = VerifyPassword(hasedPassword, password)
		assert.Nil(t, err, "Wrong Password")
	})

	t.Run("verify failed", func(t *testing.T) {
		var err = VerifyPassword(hasedPassword, "12345")
		assert.NotNil(t, err, "The password is still correct")
	})
}

// go test -v ./...
