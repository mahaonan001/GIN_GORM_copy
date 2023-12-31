package dto

import "GIN_GORM/model"

type UserJWTLogin struct {
	PassWord string `json:"password"`
	Phone    string `json:"phone"`
}
type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func JWTLogin(user model.User) UserJWTLogin {
	return UserJWTLogin{
		PassWord: user.PassWord,
		Phone:    user.Phone,
	}
}
func UserInfo(user model.User) UserDto {
	return UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
}
