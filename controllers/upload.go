package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/OlympBMSTU/annotations/config"
	"github.com/OlympBMSTU/annotations/fstorage"
	"github.com/OlympBMSTU/annotations/result"
	// "github.com/OlympBMSTU/exercises/views"
)

// UploadExerciseHandler : Controller that takes multipart form data
// parses it, saves exercise to db and sends answer to secret system
func UploadExerciseHandler(writer http.ResponseWriter, request *http.Request) {
	userID := CheckMethodAndAuthenticate(writer, request, "POST")
	if userID == nil {
		return
	}

	var err error
	if err = request.ParseMultipartForm(-1); err != nil {
		log.Print("Parse error")
		WriteResponse(&writer, "JSON", map[string]interface{}{
			"Message": "Error parse form",
			"Status":  "Error",
			"Data":    nil,
		}, http.StatusBadRequest)
		return
	}

	var fsRes result.Result
	for _, fheaders := range request.MultipartForm.File {
		for _, hdr := range fheaders {
			//	_, header, _ := request.FormFile("file")
			fsRes = fstorage.WriteFile(hdr)
			if fsRes.IsError() {
				WriteResponse(&writer, "JSON", fsRes)
				return
			}
		}
	}

	WriteResponse(&writer, "JSON", fsRes)
}
