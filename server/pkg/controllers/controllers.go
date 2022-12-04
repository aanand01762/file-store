package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aanand01762/file-store/server/pkg/libs"
	"github.com/aanand01762/file-store/server/pkg/utils"
)

type uploadResult struct {
	Filname string `json:"filename"`
	Msg     string `json:"msg"`
}

type fname struct {
	Name string `json:"filename"`
}

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

func AddFile(w http.ResponseWriter, r *http.Request) {
	var result []interface{}
	var isPartialSucess bool
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
			out := uploadResult{
				file.Filename,
				"Files uploaded successfully: " + file.Filename}
			result = append(result, out)
		} else {
			out := uploadResult{
				file.Filename,
				"file: '" + inMemFilename + "' with same content aleady exists"}
			result = append(result, out)
			isPartialSucess = true
		}
	}
	if isPartialSucess {
		w.WriteHeader(http.StatusPartialContent)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(result)

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

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	filename := &fname{}
	utils.ParseBody(r, filename)
	name := (*filename).Name

	isExist, _ := libs.CheckIfFileExists(name, "")
	if isExist {
		hash := libs.GetHashOfFile(name)
		if hash == "" {
			fmt.Fprintln(w, "Could not find the file, try updating file")
			return
		}
		libs.RemoveFile(name, hash, w)
		fmt.Fprintf(w, name+" removed successfully\n")
	} else {
		fmt.Fprintln(w, name+" does not exists in the store")
	}
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
	filenames := []string{}
	files, err := ioutil.ReadDir("./store-files/")
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	json.NewEncoder(w).Encode(filenames)
}

func GetWordCounts(w http.ResponseWriter, r *http.Request) {
	dir := "./store-files/"
	sum := libs.GetAllWordCount(dir, w)
	json.NewEncoder(w).Encode(sum)
}

/*
func GetFrequency(w http.ResponseWriter, r *http.Request) {

}
*/
