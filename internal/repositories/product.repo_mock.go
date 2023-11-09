package repositories

import (
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepoProductsMock struct {
	mock.Mock
}

func (r *RepoProductsMock) Get_Data(data *models.Products, page string, limit string, search string, orderby string) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoProductsMock) Get_Count_by_Id(id string) int {
	args := r.Mock.Called(id)
	return args.Get(0).(int)
}

func (r *RepoProductsMock) Get_Count_Data(search string) int {
	args := r.Mock.Called(search)
	return args.Get(0).(int)
}

func (r *RepoProductsMock) Insert_Data(data *models.Products) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (r *RepoProductsMock) Update_Data(data *models.Products) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (r *RepoProductsMock) Delete_Data(data *models.Products) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
