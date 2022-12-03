package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aanand01762/file-store/pkg/libs"
)

type Error struct {
	Error string `json:"Error"`
}

/*type fname struct {
	name string `json:"filename"`
}*/

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func AddFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	files := formdata.File["multiplefiles"] // grab the filenames

	for _, file := range files { // loop through the files one by one

		f, err := file.Open()
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		defer f.Close()

		byteContainer, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		hash := libs.HashFileContent(byteContainer)
		isExist, inMemFilename := libs.CheckIfFileExists(file.Filename, hash)

		if !isExist {
			libs.WriteToStore(file.Filename, hash, byteContainer, w)
			fmt.Fprintf(w, "Files uploaded successfully : ")
			fmt.Fprintf(w, file.Filename+"\n")
		} else {
			err := Error{
				"file: '" + inMemFilename + "' with same content aleady exists"}
			JSONError(w, err, 500)
		}

	}
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	defer file.Close()
	byteContainer, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	hash := libs.HashFileContent(byteContainer)
	libs.ReplaceInStore(header.Filename, hash, byteContainer, w)
}

/*
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	filename := &fname{}
	utils.ParseBody(r, filename)
	name := (*filename).name

	isExist, _ := libs.CheckIfFileExists(name, "")
	}


func GetFiles(w http.ResponseWriter, r *http.Request) {

}

func GetWordCounts(w http.ResponseWriter, r *http.Request) {

}
func GetFrequency(w http.ResponseWriter, r *http.Request) {

}
*/
