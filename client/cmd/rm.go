package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm file.txt",
	Short: "store rm file.txt",
	Long:  "store rm file.txt remove file.txt from store",
	Run: func(cmd *cobra.Command, args []string) {
		removeFiles(args)
	},
}

func removeFiles(args []string) {
	url := "http://localhost:8080/store/delete"
	method := "DELETE"
	for _, filename := range args {

		payload := strings.NewReader("{ \"filename\": \"" + filename + "\"}")
		client := &http.Client{}
		req, err := http.NewRequest(method, url, payload)

		if err != nil {
			fmt.Println(err)
			return
		}

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

}

func init() {
	rootCmd.AddCommand(rmCmd)
}
