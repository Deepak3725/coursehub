package model

import (
	"myweb/datastore/postgres"
)

type Session struct {
	Token     string `json:"token"`
	User_mail string `json:"user_mail"`
}

const CreateSession = `INSERT INTO session(token,user_mail)VALUES($1,$2)`

func (s *Session) Createsession() error {
	_, err := postgres.Db.Exec(
		CreateSession,
		s.Token,
		s.User_mail,
	)
	return err
}

const GetSes=`SELECT token,user_mail FROM session WHERE token=$1`

func(s *Session)GetSession()error{
	err:=postgres.Db.QueryRow(
		GetSes,
		s.Token,
	).Scan(&s.Token,&s.User_mail)
	return err
}

const DeleteSes=`DELETE FROM SESSION WHERE token=$1 returning token`

func(s *Session)DeleteSession()error{
	err:=postgres.Db.QueryRow(
		DeleteSes,
		s.Token,
	).Scan(&s.Token)
	return err
}
