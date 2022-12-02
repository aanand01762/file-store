package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

		out, err := os.Create("./store-files/" + file.Filename)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer out.Close()
		if err != nil {
			fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
			return
		}

		_, err = io.Copy(out, f) // file not files[i] !

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

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
