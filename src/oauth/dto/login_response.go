package oauth

import (
	userEntity "gobook/src/user/entity"
)

type LoginResponse struct {
	Token string          `json:"token"`
	User  userEntity.User `json:"user"`
}
