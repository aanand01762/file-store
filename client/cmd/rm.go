package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm file.txt",
	Short: "store rm file.txt",
	Long:  "store rm file.txt remove file.txt from store",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		removeFiles(host, port, args)
	},
}

func removeFiles(host string, port string, args []string) {

	if host == "" {
		host = viper.GetString("HOST")
	}
	if port == "" {
		port = viper.GetString("PORT")
	}

	if len(args) == 0 {
		fmt.Println("File name can not be empty")
		return
	}

	url := "http://" + host + ":" + port + "/store/delete"
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
	rmCmd.PersistentFlags().String("host", "", "File server hostname/IP")
	rmCmd.PersistentFlags().String("port", "", "File server Port number")
}
