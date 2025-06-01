package user

import "github.com/shoelfikar/finpay-models/general"

type User struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password *string `json:"password"`
	Status   string  `json:"status"`

	general.Timestamp
}
