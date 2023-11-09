package repositories

import (
	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoUsersMock struct {
	mock.Mock
}

func (rp *RepoUsersMock) Get_Users(data *models.Users, page string, limit string) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rp *RepoUsersMock) Get_Users_byId(data *models.Users) (*config.Result, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (rp *RepoUsersMock) Get_Count_by_Id(id string) int {
	args := rp.Mock.Called(id)
	return args.Get(0).(int)
}
func (rp *RepoUsersMock) Get_Count_by_Email(email string) int {
	args := rp.Mock.Called(email)
	return args.Get(0).(int)
}
func (rp *RepoUsersMock) Get_Count_by_IdEmail(email string, id string) int {
	args := rp.Mock.Called(email, id)
	return args.Get(0).(int)
}
func (rp *RepoUsersMock) Get_Count_Users() int {
	args := rp.Mock.Called()
	return args.Get(0).(int)
}
func (rp *RepoUsersMock) Insert_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (rp *RepoUsersMock) Update_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
func (rp *RepoUsersMock) Delete_User(data *models.Users) (string, error) {
	args := rp.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
