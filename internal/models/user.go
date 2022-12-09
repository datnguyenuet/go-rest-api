package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	UserID      uuid.UUID  `json:"user_id" gorm:"id" redis:"user_id" validate:"omitempty"`
	FirstName   string     `json:"first_name" gorm:"first_name" redis:"first_name" validate:"required,lte=30"`
	LastName    string     `json:"last_name" gorm:"last_name" redis:"last_name" validate:"required,lte=30"`
	Email       string     `json:"email,omitempty" gorm:"email" redis:"email" validate:"omitempty,lte=60,email"`
	Password    string     `json:"password,omitempty" gorm:"password" redis:"password" validate:"omitempty,required,gte=6"`
	Role        *string    `json:"role,omitempty" gorm:"role" redis:"role" validate:"omitempty,lte=10"`
	About       *string    `json:"about,omitempty" gorm:"about" redis:"about" validate:"omitempty,lte=1024"`
	Avatar      *string    `json:"avatar,omitempty" gorm:"avatar" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber *string    `json:"phone_number,omitempty" gorm:"phone_number" redis:"phone_number" validate:"omitempty,lte=20"`
	Address     *string    `json:"address,omitempty" gorm:"address" redis:"address" validate:"omitempty,lte=250"`
	City        *string    `json:"city,omitempty" gorm:"city" redis:"city" validate:"omitempty,lte=24"`
	Country     *string    `json:"country,omitempty" gorm:"country" redis:"country" validate:"omitempty,lte=24"`
	Gender      *string    `json:"gender,omitempty" gorm:"gender" redis:"gender" validate:"omitempty,lte=10"`
	Postcode    *int       `json:"postcode,omitempty" gorm:"postcode" redis:"postcode" validate:"omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty" gorm:"birthday" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time  `json:"created_at,omitempty" gorm:"created_at" redis:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" gorm:"updated_at" redis:"updated_at"`
	LoginDate   time.Time  `json:"login_date" gorm:"login_date" redis:"login_date"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ClearPassword() {
	u.Password = ""
}

func (u *User) PrepareUser() error {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	return nil
}

type UserList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Users      []*User `json:"users"`
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
