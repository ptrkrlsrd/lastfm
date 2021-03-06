// Copyright © 2018 Petter Karlsrud
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

// similarCmd represents the similar command
var similarCmd = &cobra.Command{
	Use:   "similar",
	Short: "Get similar artists or albums",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := client.GetSimilarArtists("swans")
		if err != nil {
			log.Fatal(err)
		}

		for _, v := range s.Artists {
			fmt.Println(v.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(similarCmd)
}
