package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string

	Ctime time.Time
	Utime time.Time
}

//func (u User) ValidationEmail() bool {
//	// Validation here
//	return u.
//}
