package model

import "myweb/datastore/postgres"

const getAllcourse = `SELECT id,title,description,price,pdf_file,image_file,author_name,bank_name,account_number,bank_qr FROM courses`

func (c *Course) ReadALL() ([]Course, error) {
	rows, erro := postgres.Db.Query(
		getAllcourse,
	)
	if erro != nil {
		return nil, erro
	}
	var courses []Course
	for rows.Next() {
		var c Course
		if err := rows.Scan(
			&c.CourseId,
			&c.Title,
			&c.Description,
			&c.Price,
			&c.Pdf_file,
			&c.Image_file,
			&c.Author_name,
			&c.Bank_Name,
			&c.Acc_Num,
			&c.Bank_Qr,
		); err != nil {
			return nil, err
		}
		courses = append(courses, c)

	}
	rows.Close()

	return courses, nil
}
