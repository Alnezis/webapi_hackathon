package manager

import (
	"api_hackathon/api"
	"api_hackathon/app"
)

type Manager struct {
	ID       int    `json:"id,omitempty" db:"id"`
	LastName string `json:"name,omitempty" db:"name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func Get(email, password string) (Manager, error) {
	i := Manager{}
	err := app.DB.Get(&i, "select id,name,email,password from users where email=? and password=?", email, password)
	api.CheckErrInfo(err, "manager get")
	return i, err
}
