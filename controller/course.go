package controller

import (
	"database/sql"
	"io"
	"myweb/model"
	httpResp "myweb/response"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCourse(w http.ResponseWriter, r *http.Request) {
	if parsErr := r.ParseMultipartForm(10 << 20); parsErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, parsErr.Error())
		return
	}

	Image_File, Image_header, Ierr := r.FormFile("image_file") //Form File opens anything sent through key image_file from frontend

	if Ierr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, "Image Upload Failed")
		return
	}
	defer Image_File.Close() //Image_File is raw bytes of image we sent...

	image_path := "./view/images/" + Image_header.Filename
	image_url := "/images/" + Image_header.Filename

	img_dst, CrErr := os.Create(image_path)

	if CrErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, CrErr.Error())
		return
	}

	defer img_dst.Close()
	io.Copy(img_dst, Image_File)
	Pdf_file, Pdf_header, Perr := r.FormFile("pdf_file")
	if Perr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, Perr.Error())
		return
	}
	defer Pdf_file.Close()
	pdf_path := "./view/pdfs/" + Pdf_header.Filename
	pdf_url := "/pdfs/" + Pdf_header.Filename

	pdf_dest, PCerr := os.Create(pdf_path)
	if PCerr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, PCerr.Error())
		return
	}
	defer pdf_dest.Close()
	io.Copy(pdf_dest, Pdf_file) //copy pdf_file in pdf_destination...

	Qr_file,Qr_header,Qr_Err:=r.FormFile("bank_qr")

	if Qr_Err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,Qr_Err.Error())
		return
	}
	defer Qr_file.Close()

	Qr_path:="./view/qrs/"+Qr_header.Filename
	Qr_Url:="/qrs/"+Qr_header.Filename

	Qr_dest,Qc_err:=os.Create(Qr_path)
	if Qc_err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,Qc_err.Error())
		return
	}
	defer Qr_dest.Close()
	io.Copy(Qr_dest,Qr_file)
	author_name := r.FormValue("author_name")
	creator_mail := r.FormValue("creator_mail")
	course := model.Course{
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		Price:        r.FormValue("price"),
		Pdf_file:     &pdf_url,
		Image_file:   &image_url,
		Author_name:  &author_name,
		Creator_mail: &creator_mail,
		Bank_Name:r.FormValue("bank_name"),
		Acc_Num: r.FormValue("acc_num"),
		Bank_Qr: &Qr_Url,
		
	}
	if CrErr := course.Create(); CrErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, CrErr.Error())
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, map[string]interface{}{ //since our id will be int so, value interface accepts both string and int...
		"Status": "Course added Successfully",
		"id":course.CourseId,  //we returned this id to front_end
	
	})
}

func GetSellingCourse(w http.ResponseWriter, r *http.Request) {
	creator_mail:=mux.Vars(r)["sid"]
	var c model.Course

	courses, readErr := c.ReadSellingList(creator_mail)
	if readErr != nil {
		switch readErr {
		case sql.ErrNoRows:
			httpResp.RespondwithError(w, http.StatusNotFound, readErr.Error())
		default:
			httpResp.RespondwithError(w, http.StatusBadRequest, readErr.Error())
		}
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, courses)
}
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	if parseErr:=r.ParseMultipartForm(10<<20);parseErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,parseErr.Error())
		return
	}

	sid := mux.Vars(r)["sid"]
	newId, idErr := getCourseId(sid)
	if idErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, idErr.Error())
		return
	}

	// 2. Check existence using a SEPARATE struct
	existingCourse := model.Course{CourseId: int(newId)}
	if existErr := existingCourse.GetCoursebyId(); existErr != nil {
		switch existErr {
		case sql.ErrNoRows:
			httpResp.RespondwithError(w, http.StatusNotFound, existErr.Error())
		default:
			httpResp.RespondwithError(w, http.StatusBadRequest, existErr.Error())
		}
		return
	}

	var imgURL string

	img_file,img_header,imgErr:=r.FormFile("image_file")

	

	if imgErr!=nil{
		switch imgErr{
		case http.ErrMissingFile:
			imgURL=*existingCourse.Image_file
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,imgErr.Error())
		    return
		}	
	}else{
		
	    defer img_file.Close()

		img_path:="./view/images/"+img_header.Filename

		imgURL="/images/"+img_header.Filename

		img_dest,Ierr:=os.Create(img_path)
		
		if Ierr!=nil{
			httpResp.RespondwithError(w,http.StatusBadRequest,Ierr.Error())
			return
		}
		defer img_dest.Close()

		io.Copy(img_dest,img_file) //dest..,file
		
	}

	var pdfUrl string

	pdf_file,pdf_header,pdf_err:=r.FormFile("pdf_file")

	if pdf_err!=nil{
		switch pdf_err{
		case http.ErrMissingFile:
			pdfUrl=*existingCourse.Pdf_file
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,pdf_err.Error())
			return
		}
	}else{
		defer pdf_file.Close()

		pdf_path:="./view/pdfs/"+pdf_header.Filename
		pdfUrl = "/pdfs/" + pdf_header.Filename
		pdf_dest,pdf_err:=os.Create(pdf_path)

		if pdf_err!=nil{
			httpResp.RespondwithError(w,http.StatusBadRequest,pdf_err.Error())
			return
		}
		defer pdf_dest.Close()

		io.Copy(pdf_dest,pdf_file)
	}
	Qr_file,Qr_header,Qr_Err:=r.FormFile("bank_qr")
	
	var Qr_Url string
	if Qr_Err!=nil{
		switch Qr_Err{
		case http.ErrMissingFile:
			Qr_Url=*existingCourse.Bank_Qr
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,Qr_Err.Error())
			return
		}
	}else{
		defer Qr_file.Close()
		Qr_path:="./view/qrs/"+Qr_header.Filename
		Qr_Url="/qrs/"+Qr_header.Filename

		Qr_dest,Qc_Err:=os.Create(Qr_path)
		if Qc_Err!=nil{
			httpResp.RespondwithError(w,http.StatusBadRequest,Qc_Err.Error())
			return
		}
		io.Copy(Qr_dest,Qr_file)
	}
	
	author_name:=r.FormValue("author_name")
	creator_mail:=r.FormValue("creator_mail")
	updated_course:=model.Course{
		Title:        r.FormValue("title"),
		Description:  r.FormValue("description"),
		Price:        r.FormValue("price"),
		Pdf_file:     &pdfUrl,
		Image_file:   &imgURL,
		Author_name:  &author_name,
		Creator_mail: &creator_mail,
		Bank_Name: r.FormValue("bank_name"),
		Acc_Num: r.FormValue("acc_num"),
		Bank_Qr: &Qr_Url,
	}
	if UpErr:=updated_course.Update(newId);UpErr!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,UpErr.Error())
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, updated_course)
}

func getCourseId(userId string) (int64, error) {
	newId, error := strconv.Atoi(userId)
	return int64(newId), error
}

func GetCoursebyId(w http.ResponseWriter,r *http.Request){
	sid:=mux.Vars(r)["sid"]
	new_id,id_Err:=getCourseId(sid)
	if id_Err!=nil{
		httpResp.RespondwithError(w,http.StatusBadRequest,id_Err.Error())
		return
	}
	c:=model.Course{CourseId: int(new_id)}
	if err:=c.GetCoursebyId();err!=nil{
		switch err{
		case sql.ErrNoRows:
			httpResp.RespondwithError(w,http.StatusNotFound,"Course Doesn't Exist")
			return
		default:
			httpResp.RespondwithError(w,http.StatusBadRequest,err.Error())
			return
		}
	}
	httpResp.RespondwithJSON(w,http.StatusOK,c)
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	sid := mux.Vars(r)["sid"]
	new_id, idErr := getCourseId(sid)
	if idErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, idErr.Error())
		return
	}
	course := model.Course{CourseId: int(new_id)}
	err := course.Delete()
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			httpResp.RespondwithError(w, http.StatusNotFound, err.Error())
		default:
			httpResp.RespondwithError(w, http.StatusBadRequest, err.Error())
		}
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, map[string]string{"Status": "Deleted"})

}
