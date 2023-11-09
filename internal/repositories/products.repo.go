package repositories

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/sasmeka/user_product_management_duazsolusi/config"
	"github.com/sasmeka/user_product_management_duazsolusi/internal/models"

	"github.com/jmoiron/sqlx"
)

type Repo_Products_IF interface {
	Get_Data(data *models.Products, page string, limit string, search string, orderby string) (*config.Result, error)
	Get_Count_by_Id(id string) int
	Get_Count_Data(search string) int
	Insert_Data(data *models.Products) (string, error)
	Update_Data(data *models.Products) (string, error)
	Delete_Data(data *models.Products) (string, error)
}

type Repo_Products struct {
	*sqlx.DB
}

func New_Products(db *sqlx.DB) *Repo_Products {
	return &Repo_Products{db}
}

func (r *Repo_Products) Get_Data(data *models.Products, page string, limit string, search string, orderby string) (*config.Result, error) {
	var list_products_data []models.Products
	var meta_product config.Metas
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

	count_data := r.Get_Count_Data(search)

	if count_data <= 0 {
		meta_product.Next = ""
	} else {
		if float64(page_int) == math.Ceil(float64(count_data)/float64(limit_int)) {
			meta_product.Next = ""
		} else {
			meta_product.Next = strconv.Itoa(page_int + 1)
		}
	}

	if page_int == 1 {
		meta_product.Prev = ""
	} else {
		meta_product.Prev = strconv.Itoa(page_int - 1)
	}

	if int(math.Ceil(float64(count_data)/float64(limit_int))) != 0 {
		meta_product.Last_page = strconv.Itoa(int(math.Ceil(float64(count_data) / float64(limit_int))))
	} else {
		meta_product.Last_page = ""
	}

	if count_data != 0 {
		meta_product.Total_data = strconv.Itoa(count_data)
	} else {
		meta_product.Total_data = ""
	}

	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}
	if orderby == "" {
		orderby = ""
	} else {
		orderby = fmt.Sprintf(` ORDER BY %s`, orderby)
	}
	q := fmt.Sprintf(`select * from products WHERE TRUE %s %s LIMIT %d OFFSET %d`, search, orderby, limit_int, offset)
	rows, err := r.Queryx(r.Rebind(q))
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		Products_data := models.Products{}
		err := rows.StructScan(&Products_data)
		if err != nil {
			log.Fatalln(err)
		}
		rows, _ := r.Queryx("select id_user, full_name, email, role from users where id_user=$1", Products_data.Id_user)
		for rows.Next() {
			Users_data := models.Users{}
			err := rows.StructScan(&Users_data)
			if err != nil {
				log.Fatalln(err)
			}
			Products_data.Detail_User = Users_data
		}
		Products_data.Id_user = nil
		list_products_data = append(list_products_data, Products_data)
	}
	rows.Close()
	if len(list_products_data) == 0 {
		return nil, errors.New("data not found.")
	}
	return &config.Result{Data: list_products_data, Meta: meta_product}, nil
}

func (r *Repo_Products) Get_Count_by_Id(id string) int {
	var count_data int
	r.Get(&count_data, "SELECT count(*) FROM public.products WHERE id_product=$1", id)
	return count_data
}

func (r *Repo_Products) Get_Count_Data(search string) int {
	if search == "" {
		search = ""
	} else {
		search = fmt.Sprintf(` AND LOWER(name_product) like LOWER('%s')`, "%"+search+"%")
	}
	var id int
	q := fmt.Sprintf(`SELECT count(*) FROM public.products WHERE TRUE %s`, search)
	r.Get(&id, r.Rebind(q))
	return id
}

func (r *Repo_Products) Insert_Data(data *models.Products) (string, error) {
	tx := r.MustBegin()

	fmt.Println(data)
	if data.Name_product == "" || data.Description == "" || data.Price == 0 || data.Stock == 0 {
		return "", errors.New("all forms must be filled")
	}

	var new_id string
	tx.Get(&new_id, "select gen_random_uuid()")
	data.Id_product = new_id
	tx.NamedExec(`INSERT INTO public.products (id_product, id_user, name_product, description, price, stock) VALUES(:id_product, :id_user, :name_product, :description, :price, :stock);`, data)
	tx.Commit()

	return "add product data successful", nil
}
func (r *Repo_Products) Update_Data(data *models.Products) (string, error) {
	tx := r.MustBegin()
	var count_data int
	tx.Get(&count_data, "SELECT count(*) FROM public.products WHERE id_product=$1 AND id_user=$2", data.Id_product, data.Id_user)
	if count_data == 0 {
		return "", errors.New("You are not allowed to update this product.")
	}
	tx.NamedExec(`UPDATE public.products SET name_product=:name_product, description=:description, stock=:stock, price=:price WHERE id_product=:id_product AND id_user = :id_user;`, data)
	tx.Commit()
	return "update product data successful", nil
}
func (r *Repo_Products) Delete_Data(data *models.Products) (string, error) {
	tx := r.MustBegin()
	var count_data int
	tx.Get(&count_data, "SELECT count(*) FROM public.products WHERE id_product=$1 AND id_user=$2", data.Id_product, data.Id_user)
	if count_data == 0 {
		return "", errors.New("You are not allowed to update this product.")
	}
	_, err := tx.NamedExec(`DELETE FROM public.products WHERE id_product=:id_product;`, data)
	if err != nil {
		return "", err
	}
	tx.Commit()
	return "delete product data successful", nil
}
