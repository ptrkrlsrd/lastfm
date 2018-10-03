// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// topCmd represents the top command
var topCmd = &cobra.Command{
	Use:   "top",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		top, err := client.GetTopTracks(username)

		if err != nil {
			log.Fatal(err)
		}

		for i, v := range top {
			fmt.Printf("%d) %s - %s\n", i+1, v.Artist.Name, v.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(topCmd)
}