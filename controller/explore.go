package controller

import (
	"myweb/model"
	httpResp "myweb/response"
	"net/http"
)

func GetAllcourse(w http.ResponseWriter, r *http.Request) {
	var course model.Course

	Courses, RlErr := course.ReadALL()
	if RlErr != nil {
		httpResp.RespondwithError(w, http.StatusBadRequest, RlErr.Error())
		return
	}
	httpResp.RespondwithJSON(w, http.StatusOK, Courses)
}
