package request

import "rest-go-stock/businesses/users"

type GetByEmailRequest struct {
	Email    string `json:"email"`
}

func ToDomainGetByemail(getbyemail GetByEmailRequest) users.Domain {
	return users.Domain{
		Email:    getbyemail.Email,
	}
}