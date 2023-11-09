package models

import (
	"time"
)

type Users struct {
	Id_user   string     `db:"id_user" json:"id_user,omitempty" form:"id_user" valid:"-"`
	Full_name string     `db:"full_name" json:"full_name" form:"full_name" valid:"required~full name is required"`
	Email     string     `db:"email" json:"email" form:"email" valid:"required~e-mail is required"`
	Pass      string     `db:"pass" json:"pass,omitempty" form:"pass" valid:"required~password is required,stringlength(6|1024)~password of at least 6 characters"`
	Role      string     `db:"role" json:"role" form:"role" valid:"-"`
	Create_at *time.Time `db:"create_at" json:"create_at,omitempty" valid:"-"`
	Update_at *time.Time `db:"update_at" json:"update_at,omitempty" valid:"-"`
}
