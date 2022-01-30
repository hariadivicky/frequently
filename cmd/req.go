/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/hariadivicky/frequently/api"
	"github.com/spf13/cobra"
)

// reqCmd represents the req command
var reqCmd = &cobra.Command{
	Use:   "req",
	Short: "make request as client to server",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			log.Fatal(err)
		}

		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatal(err)
		}

		if filePath == "" {
			log.Fatalf("file can't be blank")
		}

		max, err := cmd.Flags().GetString("max")
		if err != nil {
			log.Fatal(err)
		}

		insensitive, err := cmd.Flags().GetBool("insensitive")
		if err != nil {
			log.Fatal(err)
		}

		if err := api.MakeRequest(url, filePath, max, insensitive); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(reqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// reqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	reqCmd.Flags().BoolP("insensitive", "i", false, "insensitive mode")
	reqCmd.Flags().StringP("max", "m", "10", "max result")
	reqCmd.Flags().StringP("file", "f", "", "file to check")
	reqCmd.Flags().StringP("url", "u", "http://localhost:8000", "url for request")
}
