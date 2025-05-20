package domain

import (
	"time"
)

type AccountRole struct {
	Account_role_id int64 `gorm:"primaryKey"`
	Account_id      int64
	Role_id         int64
	Create_at		*time.Time
}

type AccountRoleRepository interface {
	Create(*AccountRole) (*AccountRole, error)
	FindByAccountId(int64) ([]*AccountRole, error)
}