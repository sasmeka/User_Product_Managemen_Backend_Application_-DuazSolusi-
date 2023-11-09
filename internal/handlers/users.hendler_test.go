package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	_ "image/png"
)

var repoUsersMock = repositories.RepoUsersMock{}
var reqBody = `{
	"full_name":"sasmeka",
	"email":"verdysas@gmail.com",
	"pass":"123456"
}`

func TestGet_Data_Users(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Users(&repoUsersMock)
	expect := &config.Result{
		Data: []interface{}{map[string]interface{}{
			"id_user":   "c458f82c-54b2-4125-90db-f5a9b817ac02",
			"full_name": "verdi",
			"email":     "sas2@gmail.com",
			"role":      "user",
			"create_at": "2023-11-08T23:50:50.808067Z",
			"update_at": interface{}(nil)}},
		Meta: map[string]interface{}{"last_page": "1", "next": "", "prev": "", "total_data": "1"}}
	repoUsersMock.On("Get_Users", mock.Anything).Return(expect, nil)

	r.GET("/user", handler.Get_Data_Users)
	req := httptest.NewRequest("GET", "/user?limit=1&page=1", nil)
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":200, "data": [
        {
			"id_user": "c458f82c-54b2-4125-90db-f5a9b817ac02",
            "full_name": "verdi",
            "email": "sas2@gmail.com",
            "role": "user",
            "create_at": "2023-11-08T23:50:50.808067Z",
            "update_at": null
		}
    ], "meta": {
        "next": "",
        "prev": "",
        "last_page": "1",
        "total_data": "1"
    } ,"status":"OK"}`, w.Body.String())
}

func TestPost_Data_User(t *testing.T) {
	t.Run("post success", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		_, r := gin.CreateTestContext(w)

		handler := New_Users(&repoUsersMock)
		repoUsersMock.On("Get_Count_by_Email", mock.Anything).Return(0)
		repoUsersMock.On("Insert_User", mock.Anything).Return("add user data successful.", nil)

		r.POST("/create_user", handler.Post_Data_User)
		req := httptest.NewRequest("POST", "/create_user", strings.NewReader(reqBody))
		req.Header.Set("Content-type", "application/json")
		r.ServeHTTP(w, req)

		if w.Code == 400 {
			assert.Equal(t, 400, w.Code)
			assert.JSONEq(t, `{"code":400, "description": "e-mail already registered.", "status":"Bad Request"}`, w.Body.String())
		} else {
			assert.Equal(t, http.StatusOK, w.Code)
			assert.JSONEq(t, `{"code":200, "description": "add user data successful.", "status":"OK"}`, w.Body.String())
		}
	})
}

func TestPut_Data_User(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Users(&repoUsersMock)
	count_id_put := 1
	repoUsersMock.On("Get_Count_by_Id", mock.Anything).Return(count_id_put)
	repoUsersMock.On("Get_Count_by_IdEmail", mock.Anything, mock.Anything).Return(0)
	repoUsersMock.On("Update_User", mock.Anything).Return("update user data successful", nil)

	r.PUT("/update_user/:id", handler.Put_Data_User)
	req := httptest.NewRequest("PUT", "/update_user/asdg8awgd8wtd6", strings.NewReader(reqBody))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "update user data successful", "status":"OK"}`, w.Body.String())
	}
}

func TestDelete_Data_User(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Users(&repoUsersMock)
	count_id_del := 1
	repoUsersMock.On("Get_Count_by_Id", mock.Anything).Return(count_id_del)
	repoUsersMock.On("Delete_User", mock.Anything).Return("delete user data successful", nil)

	r.DELETE("/delete_user/:id", handler.Delete_Data_User)
	req := httptest.NewRequest("DELETE", "/delete_user/asdg8awgd8wtd6", strings.NewReader("{}"))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "delete user data successful", "status":"OK"}`, w.Body.String())
	}
}
