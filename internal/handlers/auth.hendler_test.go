package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"
	"github.com/sasmeka/user_product_management_duazsolusi/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoAuthMock = repositories.RepoAuthMock{}
var reqBodyReg = `{
    "email":"verdi@gmail.com",
    "pass":"123456",
    "full_name":"verdi"
}`
var reqBodyLogin = `{
    "email":"verdi@gmail.com",
    "pass":"123456"
}`

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Auth(&repoAuthMock)
	expect := &models.Users{
		Id_user:   "2067d582-1013-4b91-b0bd-b13a968a7b93",
		Full_name: "verdi",
		Email:     "verdi@gmail.com",
		Role:      "user",
		Pass:      "$2a$10$ejKXeNiZ4dmHxBDPBzcoU.asRezCXFNZtmf1gnpRfX8ruTP9eDRP6",
	}
	repoAuthMock.On("Get_User", mock.Anything).Return(expect, nil)
	jwtt := pkg.NewToken(expect.Id_user, expect.Role, expect.Email)
	tokens, _ := jwtt.Generate()

	r.POST("/login", handler.Login)
	req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBodyLogin))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "wrong password", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "status":"OK", "token":"`+tokens+`"}`, w.Body.String())
	}
}

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Auth(&repoAuthMock)
	repoAuthMock.On("Get_Count_by_Email", mock.Anything).Return(0)
	repoAuthMock.On("Register_rep", mock.Anything).Return("register successful.", nil)

	r.POST("/register", handler.Register)
	req := httptest.NewRequest("POST", "/register", strings.NewReader(reqBodyReg))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "e-mail already registered.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "register successful.", "status":"OK"}`, w.Body.String())
	}
}
