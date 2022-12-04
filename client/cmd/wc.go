package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc",
	Short: "store wc",
	Long: `store wc returns the number of words 
	in all the files stored in server`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		wordCount(host, port)
	},
}

func wordCount(host string, port string) {
	if host == "" {
		host = viper.GetString("HOST")
	}
	if port == "" {
		port = viper.GetString("PORT")
	}

	url := "http://" + host + ":" + port + "/store/word-count"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
func init() {
	rootCmd.AddCommand(wcCmd)
	wcCmd.PersistentFlags().String("host", "", "File server hostname/IP")
	wcCmd.PersistentFlags().String("port", "", "File server Port number")
}
