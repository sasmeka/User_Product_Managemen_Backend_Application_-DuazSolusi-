package repositories

import (
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
	"github.com/stretchr/testify/mock"
)

type RepoAuthMock struct {
	mock.Mock
}

func (r *RepoAuthMock) Get_Count_by_Email(email string) int {
	args := r.Mock.Called(email)
	return args.Get(0).(int)
}
func (r *RepoAuthMock) Get_User(data *models.Auth) (*models.Users, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*models.Users), args.Error(1)
}
func (r *RepoAuthMock) Register_rep(data *models.Users) (string, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(string), args.Error(1)
}
