package dto

import "ginEssential/model"

/*
	feature:处理输出流，隐藏关键信息。
 */
type UserDto struct {
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user model.User)UserDto{
	return UserDto{
		Name: user.Name,
		Telephone: user.Telephone,
	}
}