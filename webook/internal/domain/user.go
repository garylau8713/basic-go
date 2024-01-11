package domain

import "time"

type User struct {
	Id       int64
	Email    string
	Password string

	// Use UTC 0 time
	Ctime time.Time
	Utime time.Time
}

//func (u User) ValidationEmail() bool {
//	// Validation here
//	return u.
//}
