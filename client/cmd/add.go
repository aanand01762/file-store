package cmd

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <file1> <file2> ",
	Short: "store add <file1> <file2> ",
	Long: `store add file1.txt file2.txt send 
	both files - file1.txt and file2.txt in the 
	current path to the file store. Add command should 
	fail if the file already exists in the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		uploadfiles(args)
	},
}

func uploadfiles(args []string) {

	url := "http://localhost:8080/store/add"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	for _, path := range args {

		file, err := os.Open(path)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()
		part1, errFile := writer.CreateFormFile("multiplefiles", filepath.Base(path))
		if errFile != nil {
			fmt.Println(errFile)
			return
		}

		_, errFile = io.Copy(part1, file)
		if errFile != nil {
			fmt.Println(errFile)
			return
		}
	}

	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func init() {
	rootCmd.AddCommand(addCmd)

}
