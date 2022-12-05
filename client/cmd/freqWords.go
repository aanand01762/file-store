package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// freqWordsCmd represents the freqWords command
var freqWordsCmd = &cobra.Command{
	Use:   "freq-words",
	Short: "store freq-words [--limit|-n 10] [--order=dsc|asc]",
	Long: `Store freq-words return the least or most frequent
	words in all the files combined. By default, command 
	return the 10 most frequent words in all the files combined.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		limit, _ := cmd.Flags().GetInt("limit")
		order, _ := cmd.Flags().GetString("order")
		getFrequency(host, port, limit, order, args)
	},
}

func getFrequency(host string, port string, limit int, order string, args []string) {
	type input struct {
		Order string `json:"order"`
		Limit int    `json:"limit"`
	}

	data := input{
		Order: order,
		Limit: limit,
	}

	payloadBytes, err := json.Marshal(data)

	payload := bytes.NewBuffer(payloadBytes)
	if host == "" {
		host = viper.GetString("HOST")
	}
	if port == "" {
		port = viper.GetString("PORT")
	}
	if len(args) > 0 {
		fmt.Printf("Worng args passed " + strings.Join(args, " "))
		return
	}

	url := "http://" + host + ":" + port + "/store/frequency"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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
	rootCmd.AddCommand(freqWordsCmd)
	freqWordsCmd.PersistentFlags().String("host", "", "File server hostname/IP")
	freqWordsCmd.PersistentFlags().String("port", "", "File server Port number")
	freqWordsCmd.PersistentFlags().Int("limit", 10, "Number of words to display")
	freqWordsCmd.PersistentFlags().String("order", "dsc", "order of frequency: asc|dsc")
}
