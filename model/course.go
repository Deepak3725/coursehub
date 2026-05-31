package model

import (
	"log"
	"myweb/datastore/postgres"
)

type Course struct {
	CourseId    int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string    `json:"price"`
	Pdf_file    *string `json:"pdf_file"`
	Image_file  *string `json:"image_file"`
	Author_name *string `json:"author_name"` 
	Creator_mail *string `json:"creator_mail"`
	Bank_Name  string `json:"bank_name"`
	Acc_Num string `json:"acc_num"`
	Bank_Qr *string `json:"bank_qr"`
}

const courseInsert = `INSERT INTO courses(title,description,price,pdf_file,image_file,author_name,creator_mail,bank_name,account_number,bank_qr)
VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`

func (c *Course) Create() error {
	 err := postgres.Db.QueryRow( //we are using QueryRow instead of Exec because we want to insert and get that row back
		courseInsert,
		c.Title,
		c.Description,
		c.Price,
		c.Pdf_file,
		c.Image_file,
		c.Author_name,
		c.Creator_mail,
		c.Bank_Name,
		c.Acc_Num,
		c.Bank_Qr,
	).Scan(&c.CourseId) //we capture the course_id here returned by db
	return err
}
//This is for search//
const getSellingCourse = `SELECT id,title,description,price,pdf_file,image_file,author_name,bank_name,account_number,bank_qr FROM courses WHERE creator_mail=$1`

func (c *Course) ReadSellingList(creator_mail string) ([]Course, error) {
	rows, err := postgres.Db.Query(
		getSellingCourse,
		creator_mail,
	)
	if err != nil {
		return nil, err
	}

	var courses []Course = []Course{}  // empty slice, not nil

	for rows.Next() {
		var c Course
		if erro := rows.Scan(
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
		); erro != nil {
			return nil, erro
		}
		courses = append(courses, c)
	}

	rows.Close()

	return courses, err

}



const updateQuery = `UPDATE courses SET 
    title = COALESCE(NULLIF($1, ''), title),
    description = COALESCE(NULLIF($2, ''), description),
    price = COALESCE(NULLIF($3, ''), price),
    pdf_file = COALESCE(NULLIF($4, ''), pdf_file),
    image_file = COALESCE(NULLIF($5, ''), image_file),
    author_name = COALESCE(NULLIF($6, ''), author_name),
	bank_name = COALESCE(NULLIF($7, ''), bank_name),
	account_number = COALESCE(NULLIF($8, ''), account_number),
	bank_qr = COALESCE(NULLIF($9, ''), bank_qr)
    WHERE id = $10`

func (c *Course) Update(oldId int64) error {
	result, err := postgres.Db.Exec(
		updateQuery,
		c.Title,
		c.Description,
		c.Price,
		c.Pdf_file,
		c.Image_file,
		c.Author_name,
		c.Bank_Name,
		c.Acc_Num,
		c.Bank_Qr,
		oldId,
	)
	if err != nil {
		log.Println("Update error:", err)
		return err
	}

	rows, _ := result.RowsAffected()
	log.Println("Rows affected:", rows)

	return nil
}

const getcoursebyid = `SELECT title,description,price,pdf_file,image_file,author_name,creator_mail,bank_name,account_number,bank_qr FROM courses WHERE id=$1`

func (c *Course) GetCoursebyId() error {
	err := postgres.Db.QueryRow(
		getcoursebyid,
		c.CourseId,
	).Scan(
		&c.Title,
		&c.Description,
		&c.Price,
		&c.Pdf_file,
		&c.Image_file,
		&c.Author_name,
		&c.Creator_mail,
		&c.Bank_Name,
		&c.Acc_Num,
		&c.Bank_Qr,
	)
	return err
}

const deleteCourse = `DELETE FROM courses WHERE id=$1 RETURNING id,title`

func (c *Course) Delete() error {
	err := postgres.Db.QueryRow(
		deleteCourse,
		c.CourseId,
	).Scan(&c.CourseId, &c.Title)
	return err
}
