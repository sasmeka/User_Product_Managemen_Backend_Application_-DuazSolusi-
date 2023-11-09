package repositories

import (
	"errors"
	"math"
	"strconv"

	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Users_IF interface {
	Get_Users(data *models.Users, page string, limit string) (*config.Result, error)
	Get_Users_byId(data *models.Users) (*config.Result, error)
	Get_Count_by_Id(id string) int
	Get_Count_by_Email(email string) int
	Get_Count_by_IdEmail(email string, id string) int
	Get_Count_Users() int
	Insert_User(data *models.Users) (string, error)
	Update_User(data *models.Users) (string, error)
	Delete_User(data *models.Users) (string, error)
}

type Repo_Users struct {
	*sqlx.DB
}

func New_Users(db *sqlx.DB) *Repo_Users {
	return &Repo_Users{db}
}

func (r *Repo_Users) Get_Users(data *models.Users, page string, limit string) (*config.Result, error) {
	users_data := []models.Users{}

	var meta_user config.Metas
	var offset int = 0
	var page_int, _ = strconv.Atoi(page)
	var limit_int, _ = strconv.Atoi(limit)
	if limit == "" {
		limit_int = 5
	}
	if page == "" {
		page_int = 1
	}
	if page_int > 0 {
		offset = (page_int - 1) * limit_int
	} else {
		offset = 0
	}

	count_data := r.Get_Count_Users()

	if count_data <= 0 {
		meta_user.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_user.Next = ""
		} else {
			meta_user.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_user.Prev = ""
	} else {
		meta_user.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_user.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_user.Last_page = ""
	}

	if count_data != 0 {
		meta_user.Total_data = strconv.Itoa(count_data)
	} else {
		meta_user.Total_data = ""
	}
	r.Select(&users_data, `SELECT id_user,full_name, email, "role", create_at, update_at FROM public.users LIMIT $1 OFFSET $2`, limit_int, offset)
	if len(users_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: users_data, Meta: meta_user}, nil
}

func (r *Repo_Users) Get_Users_byId(data *models.Users) (*config.Result, error) {
	users_data := []models.Users{}
	r.Select(&users_data, `SELECT id_user,full_name, email, "role", create_at, update_at FROM public.users WHERE id_user=$1`, data.Id_user)
	if len(users_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: users_data}, nil
}

func (r *Repo_Users) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE id_user=$1", id)
	return count_data
}

func (r *Repo_Users) Get_Count_by_Email(email string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE LOWER(email)=LOWER($1)", email)
	return count_data
}

func (r *Repo_Users) Get_Count_by_IdEmail(email string, id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.users WHERE LOWER(email)=LOWER($1) AND id_user!=$2", email, id)
	return count_data
}

func (r *Repo_Users) Get_Count_Users() int {
	var id int
	r.Get(&id, "SELECT count(*) FROM public.users")
	return id
}

func (r *Repo_Users) Insert_User(data *models.Users) (string, error) {
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
	return "add user data successful.", nil
}
func (r *Repo_Users) Update_User(data *models.Users) (string, error) {
	query := `UPDATE public.users SET
			full_name=:full_name, 
			email=:email, 
			pass=:pass,
			update_at=now()
			WHERE id_user=:id_user;`
	if data.Full_name == "" || data.Email == "" || data.Pass == "" {
		return "", errors.New("all forms must be filled")
	}
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "update user data successful", nil
}
func (r *Repo_Users) Delete_User(data *models.Users) (string, error) {
	query := `DELETE FROM public.users WHERE id_user=:id_user;`
	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}
	return "delete user data successful", nil
}
