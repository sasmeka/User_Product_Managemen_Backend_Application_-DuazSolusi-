package pkg

import (
	"testing"

	"github.com/sasmeka/user_product_management_duazsolusi/config"

	"github.com/stretchr/testify/assert"
)

var status_code = 500
var status_code_fail = 1000
var desc = ""

func TestGetStatus(t *testing.T) {
	t.Run("check desc success", func(t *testing.T) {
		var response = getStatus(status_code)
		desc = response
		assert.NotEqual(t, "", response, "No description")
	})

	t.Run("check desc failed", func(t *testing.T) {
		var response = getStatus(status_code_fail)
		assert.Equal(t, "", response, "Description is still there")
	})
}

func TestResponse(t *testing.T) {
	t.Run("check response", func(t *testing.T) {
		expected := Response{Code: status_code, Status: desc}
		var response = Responses(status_code, &config.Result{})
		assert.Equal(t, expected, *response, "Response not equal")
	})
}
