package pkg

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var token = ""
var uuid = "1234567890"
var role = "admin"
var email = "verdysas@gmail.com"
var data_token = claims{}

func TestNewToken(t *testing.T) {
	t.Run("Set data token", func(t *testing.T) {
		var expected = claims{Id: uuid, Role: role, Email: email, RegisteredClaims: jwt.RegisteredClaims{Issuer: "product_duazsolusi"}}
		var response = NewToken(uuid, role, email)
		response.RegisteredClaims.ExpiresAt = nil
		data_token = *response
		assert.Equal(t, expected, *response, "Data token results do not match")
	})
}

func TestGenerate(t *testing.T) {
	t.Run("genrate token", func(t *testing.T) {
		response, err := data_token.Generate()
		token = response
		assert.NotEqual(t, "", response, "Empty token")
		assert.Nil(t, err, "Generating token error")
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("verify token success", func(t *testing.T) {
		var expected = data_token
		response, err := VerifyToken(token)
		assert.Equal(t, expected, *response, "Data token results do not match")
		assert.Nil(t, err, "Generating token error")
	})
}
