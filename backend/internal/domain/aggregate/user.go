
package aggregate

import "time"

type User struct {
	ID       string
	Email    string
	Username string
	Password string
	Created  time.Time
	Updated  time.Time
}

func NewUser(id, email, username, password string) *User {
	return &User{
		ID:       id,
		Email:    email,
		Username: username,
		Password: password,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
}
