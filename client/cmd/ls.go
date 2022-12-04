/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "store ls",
	Long:  "store ls list all files in the store",
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")
		listFiles(host, port)
	},
}

func listFiles(host string, port string) {

	if host == "" {
		host = viper.GetString("HOST")
	}
	if port == "" {
		port = viper.GetString("PORT")
	}

	url := "http://" + host + ":" + port + "/store/files"
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
	var data []string
	json.Unmarshal(body, &data)
	for _, val := range data {
		fmt.Printf("%s ", val)
	}
	fmt.Printf("\n")
}

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.PersistentFlags().String("host", "", "File server hostname/IP")
	lsCmd.PersistentFlags().String("port", "", "File server Port number")
}
