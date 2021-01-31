package model

import "time"

type User struct {
	Id int `json:"id"`
	Email string `json:"email"`
	UserName string `json:"username" gorm:"column:username"`
	Password string `json:"password"`
	Status int `json:"status"`
	IsDel int `json:"is_del"`
	CreatedAt *time.Time `json:"created_at" gorm:"-"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"-"`
}

func (User) TableName() string  {
	return "user"
}