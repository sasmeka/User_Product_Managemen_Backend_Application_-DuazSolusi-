package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoProductsMock = repositories.RepoProductsMock{}
var reqBodyProduct = `{
    "name_product":"coba",
    "description":"pembersih",
    "price":5000,
    "stock":10
}`

func TestGet_Data_Products(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Products(&repoProductsMock)
	expect := &config.Result{
		Data: []interface{}{map[string]interface{}{
			"id_product":   "17b577aa-323f-49c2-a2ba-231d1a9794e5",
			"name_product": "coba",
			"description":  "pembersih",
			"price":        5000,
			"stock":        10,
			"create_at":    "2023-11-09T09:52:35.343085Z",
			"update_at":    interface{}(nil),
			"detail_users": map[string]interface{}{
				"id_user":   "2067d582-1013-4b91-b0bd-b13a968a7b93",
				"full_name": "verdi",
				"email":     "verdi@gmail.com",
				"role":      "user"}}},
		Meta: map[string]interface{}{"last_page": "1", "next": "", "prev": "", "total_data": "1"}}
	repoProductsMock.On("Get_Data", mock.Anything).Return(expect, nil)

	r.GET("/product", handler.Get_Data_Products)
	req := httptest.NewRequest("GET", "/product?limit=1&page=1", nil)
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"code":200, "data": [
        {
            "id_product": "17b577aa-323f-49c2-a2ba-231d1a9794e5",
            "name_product": "coba",
            "description": "pembersih",
            "price": 5000,
            "stock": 10,
            "create_at": "2023-11-09T09:52:35.343085Z",
            "update_at": null,
            "detail_users": {
                "id_user": "2067d582-1013-4b91-b0bd-b13a968a7b93",
                "full_name": "verdi",
                "email": "verdi@gmail.com",
                "role": "user"
            }
        }
    ], "meta": {
        "next": "",
        "prev": "",
        "last_page": "1",
        "total_data": "1"
    } ,"status":"OK"}`, w.Body.String())
}

func TestPost_Data_Product(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Products(&repoProductsMock)
	repoProductsMock.On("Insert_Data", mock.Anything).Return("add product data successful.", nil)

	r.POST("/product", handler.Post_Data_Product)
	req := httptest.NewRequest("POST", "/product", strings.NewReader(reqBodyProduct))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "all forms must be filled", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "add product data successful.", "status":"OK"}`, w.Body.String())
	}
}

func TestPut_Data_Product(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Products(&repoProductsMock)
	count_id_put := 1
	repoProductsMock.On("Get_Count_by_Id", mock.Anything).Return(count_id_put)
	repoProductsMock.On("Update_Data", mock.Anything).Return("update product data successful", nil)

	r.PUT("/product/:id", handler.Put_Data_Product)
	req := httptest.NewRequest("PUT", "/product/asdg8awgd8wtd6", strings.NewReader(reqBodyProduct))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "update product data successful", "status":"OK"}`, w.Body.String())
	}
}

func TestDelete_Data_Product(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	handler := New_Products(&repoProductsMock)
	count_id_put := 1
	repoProductsMock.On("Get_Count_by_Id", mock.Anything).Return(count_id_put)
	repoProductsMock.On("Delete_Data", mock.Anything).Return("delete product data successful", nil)

	r.DELETE("/product/:id", handler.Delete_Data_Product)
	req := httptest.NewRequest("DELETE", "/product/asdg8awgd8wtd6", strings.NewReader(reqBodyProduct))
	req.Header.Set("Content-type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code == 400 {
		assert.Equal(t, 400, w.Code)
		assert.JSONEq(t, `{"code":400, "description": "data not found.", "status":"Bad Request"}`, w.Body.String())
	} else {
		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"code":200, "description": "delete product data successful", "status":"OK"}`, w.Body.String())
	}
}
