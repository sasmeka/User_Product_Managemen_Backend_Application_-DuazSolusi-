package models

import (
	"time"
)

type Products struct {
	Id_product   string     `json:"id_product" db:"id_product" form:"id_product" valid:"-"`
	Name_product string     `json:"name_product" db:"name_product" form:"name_product" valid:"required~name product is required"`
	Description  string     `json:"description" db:"description" form:"description" valid:"required~description is required"`
	Price        int        `json:"price" db:"price" form:"price" valid:"required~price is required"`
	Stock        int        `json:"stock" db:"stock" valid:"required~stock is required"`
	Create_at    *time.Time `json:"create_at" db:"create_at" form:"create_at" valid:"-"`
	Update_at    *time.Time `json:"update_at" db:"update_at" form:"update_at" valid:"-"`
	Id_user      *string    `json:"id_user,omitempty" db:"id_user" form:"id_user" valid:"-"`
	Detail_User  Users      `json:"detail_users,omitempty" valid:"-"`
}
