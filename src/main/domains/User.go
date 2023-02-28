package main_domains

import "time"

type User struct {
	Id          string
	Name        string
	Nick        string
	Email       string
	Password    string
	CreatedDate time.Time
}
