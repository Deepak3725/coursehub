package model

import "myweb/datastore/postgres"

type Admin struct {
	FirstName string `json:"fname"`
	LastName  string `json:"lname"`
	Password  string `json:"pw"`
	Email     string `json:"email"`
	Role      string `json:"role"`
}

const signUpQuery = `INSERT INTO admin(firstname,lastname,password,email,role)VALUES($1,$2,$3,$4,$5)`

func (adm *Admin) SignUp() error {
	_,err:=postgres.Db.Exec(
		signUpQuery,
		adm.FirstName,
		adm.LastName,
		adm.Password,
		adm.Email,
		adm.Role,
	)
	return err
}

const LogInQuery=`SELECT email,password FROM admin WHERE email=$1 and password=$2`

func(adm *Admin)LogIn()error{
	err:=postgres.Db.QueryRow(
		LogInQuery,
		adm.Email,
		adm.Password,
	).Scan(&adm.Email,&adm.Password)
	return err
}


