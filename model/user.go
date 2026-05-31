package model

import "myweb/datastore/postgres"

type User struct {
	UserId   int    `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

const insertquery = `INSERT INTO users(id,name,email,password,role)
VALUES($1,$2,$3,$4,$5)`

func (u *User) Create() error {
	_, err := postgres.Db.Exec(
		insertquery,
		u.UserId,
		u.UserName,
		u.Email,
		u.Password,
		u.Role,
	)
	return err
}

const readuser=`SELECT id,name,email,role FROM users WHERE id=$1`

func (u *User) Read()error{
	err:=postgres.Db.QueryRow(
		readuser,
		u.UserId,
	).Scan(
		&u.UserId,
		&u.UserName,
		&u.Email,
		&u.Role,
	)
	return err
}

