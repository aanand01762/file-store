package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aanand01762/file-store/pkg/libs"
)

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
		fmt.Fprintf(w, hash+"\n")

		filename := "./store-files/" + file.Filename
		libs.WriteToStore(filename, byteContainer, w)

		fmt.Fprintf(w, "Files uploaded successfully : ")
		fmt.Fprintf(w, file.Filename+"\n")

	}
}

/*
func DeleteFile(w http.ResponseWriter, r *http.Request) {


}

func UpdateFile(w http.ResponseWriter, r *http.Request) {

	}


func GetFiles(w http.ResponseWriter, r *http.Request) {

}

func GetWordCounts(w http.ResponseWriter, r *http.Request) {

}
func GetFrequency(w http.ResponseWriter, r *http.Request) {

}
*/
