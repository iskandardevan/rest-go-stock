package request

import "rest-go-stock/businesses/users"

type RegisterUserRequest struct {
	Email    	string 	`json:"email"`
	Name     	string 	`json:"name"` 
	Password 	string 	`json:"password"`
	Roles_ID	uint	`json:"roles_id"`
}

func (User *RegisterUserRequest) ToDomain() *users.Domain {
	return &users.Domain{
		Email	:User.Email,
		Name    :User.Name, 
		Password:User.Password, 
	}
}