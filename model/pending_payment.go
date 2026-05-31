package model

import "myweb/datastore/postgres"

type Payment struct{
	Token string `json:"token"`
	Buyer_mail string `json:"buyer_mail"`
	CourseId int `json:"course_id"`
	Reference_Num string `json:"ref_num"`

}

const insertPayment=`INSERT INTO pending_payments(token,buyer_mail,course_id,reference_num) VALUES($1,$2,$3,$4)`

func(p *Payment)CreatePay()error{
	_,err:=postgres.Db.Exec(
		insertPayment,
		p.Token,
		p.Buyer_mail,
		p.CourseId,
		p.Reference_Num,
	)
	return err
}

const getPayment=`SELECT token,buyer_mail,course_id,reference_num FROM pending_payments WHERE token=$1`

func (p *Payment)ReadPayment()error{
	err:=postgres.Db.QueryRow(
		getPayment,
		p.Token,
	).Scan(
		&p.Token,
		&p.Buyer_mail,
		&p.CourseId,
		&p.Reference_Num,
	)
	return err
}