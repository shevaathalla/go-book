package user

import "database/sql"

type User struct {
	ID              int64        `gorm:"primaryKey" json:"id"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Password        string       `json:"-"`
	CodeVerified    string       `json:"code_verified"`
	EmailVerifiedAt sql.NullTime `json:"email_verified_at"`
	Instance        string       `json:"instance"`
	Address         string       `json:"address"`
	Phone           string       `json:"phone"`
	CreatedAt       sql.NullTime `json:"created_at"`
	UpdatedAt       sql.NullTime `json:"updated_at"`
}
