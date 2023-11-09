package repositories

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"
)

type Repo_AuthIF interface {
	Get_Count_by_Email(email string) int
	Get_User(data *models.Auth) (*models.Users, error)
	Register_rep(data *models.Users) (string, error)
}

type Repo_Auth struct {
	*sqlx.DB
}

func New_Auth(db *sqlx.DB) *Repo_Auth {
	return &Repo_Auth{db}
}

func (r *Repo_Auth) Get_Count_by_Email(email string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE LOWER(email)=LOWER($1)", email)
	return count_data
}

func (r *Repo_Auth) Get_User(data *models.Auth) (*models.Users, error) {
	var result models.Users

	q := `SELECT id_user, email, "role", "pass" FROM public.users WHERE email = ?`

	if err := r.Get(&result, r.Rebind(q), data.Email); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("email not found")
		}

		return nil, err
	}
	return &result, nil
}
func (r *Repo_Auth) Register_rep(data *models.Users) (string, error) {
	query := `INSERT INTO public.users(
		full_name,
		email, 
		pass
	)VALUES(
		:full_name,
		:email, 
		:pass
	);`
	if data.Full_name == "" || data.Email == "" || data.Pass == "" {
		return "", errors.New("all forms must be filled")
	}
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "register successful.", nil
}
