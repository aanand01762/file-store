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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "store update file.txt",
	Long: `store update file.txt update contents of 
	file.txt in server with the local file.txt or 
	create a new file.txt in server if it is absent.`,
	Run: func(cmd *cobra.Command, args []string) {
		updateFile(args)
	},
}

func updateFile(args []string) {
	url := "http://localhost:8080/store/update"
	method := "PUT"

	if len(args) > 1 {
		fmt.Println("Passing multiple files args not supported")
		return
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	path := args[0]
	file, errFile1 := os.Open(path)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(path))
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}

	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
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
	rootCmd.AddCommand(updateCmd)
}
