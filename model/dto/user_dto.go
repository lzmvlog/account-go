package dto

import "account-go/model"

type UserDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// ToUserDTO 转换用户 dto
func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Id:   user.Id,
		Name: user.UserName,
	}
}
