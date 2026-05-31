package model

import (
	"myweb/datastore/postgres"
)

type Purchase struct {
	Purchase_id int    `json:"id"`
	Course_id   int    `json:"course_id"`
	User_email  string `json:"user_email"`
}

const purchaseInsert = `INSERT INTO purchases(course_id,user_email)
VALUES($1,$2)`

func (p *Purchase) Create() error {
	_, err := postgres.Db.Exec(
		purchaseInsert,
		p.Course_id,
		p.User_email,
	)
	return err
}

type PurchasedCourse struct {
	CourseId int `json:"c_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image_file  *string `json:"image_file"`
	Pdf_file    *string `json:"pdf_file"`
	Author_name *string `json:"author_name"`
}

const getPurchasequery = `SELECT c.id,c.title,c.description,c.image_file,c.pdf_file,c.author_name FROM purchases p
JOIN courses c ON p.course_id=c.id WHERE p.user_email=$1`

func (p *PurchasedCourse) ReadPurchasedAll(user_email string) ([]PurchasedCourse, error) {
	rows, err := postgres.Db.Query(
		getPurchasequery,
		user_email,
	)
	if err != nil {
		return nil, err
	}
var purchased []PurchasedCourse = []PurchasedCourse{}  //  empty slice, not nil
	for rows.Next() {
		var p PurchasedCourse
		if erro := rows.Scan(
			&p.CourseId,
			&p.Title,
			&p.Description,
			&p.Image_file,
			&p.Pdf_file,
			&p.Author_name,
		); erro != nil {
			return nil, erro

		}
		purchased = append(purchased, p)

	}
	rows.Close()
	return purchased, nil
}
