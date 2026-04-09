package db

type UserTable struct {
	Id       uint64
	Gmail    string
	Password string
}

var Users []UserTable
