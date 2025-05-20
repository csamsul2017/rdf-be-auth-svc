package domain

import (
	"time"
)

type Account struct {
	Id				int64		`gorm:"primaryKey"`
	Username 		string		`gorm:"unique:not null"`
	Email			string		`gorm:"unique;not null"`
	Password_hash	string		`gorm:"not null"`
	Is_active 		bool
	Create_at		*time.Time
	Update_at		*time.Time
	Delete_at		*time.Time
}

type CreateAccountRequest struct {
	Username		string 		`validate:"required"`
	Email			string 		`validate:"required,email"`
	Password_hash 	string 		`validate:"required,min=8"`
	Role			string 		`validate:"required"`		
}

type LoginRequest struct {
	Username		string 		`validate:"required"`
	Password		string 		`validate:"required"`
}

type UpdateAccountRequest struct {
	Id				int64		`validate:"required"`
	Email			string		`validate:"omitempty,email"`
	Password		string		`validate:"omitempty,min=8"`
	Role     		string 		`validate:"omitempty"`
	IsActive 		*bool
}

type DeleteAccountRequest struct {
	Id 				int64		`validate:"required"`
}

type AccountRepository interface {
	Create(*Account) (*Account, error)
	FindByID(id int64) (*Account, error)
	FindByUsername(username string) (*Account, error)
	Update(*Account) (*Account, error)
	Delete(*Account) error
}

type AccountUsecase interface {
	Register(*CreateAccountRequest) (*Account, error)
	Login(*LoginRequest) (*Account, error)
	GetAccount(id int64) (*Account, error)
	UpdateAccount(*UpdateAccountRequest) (*Account, error)
	DeleteAccount(id int64) error
}