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

	"github.com/spf13/cobra"
)

var limit = 10

// scrobblesCmd represents the scrobbles command
var scrobblesCmd = &cobra.Command{
	Use:   "scrobbles",
	Short: "Get an users top scrobbles",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		topTracks, err := client.GetRecentTracks(username)
		if err != nil {
			handleError(err)
		}

		if len(topTracks.Tracks) < limit {
			limit = len(topTracks.Tracks)
		}

		for _, v := range topTracks.Tracks[:limit] {
			fmt.Printf(v.ToString())
		}
	},
}

func init() {
	scrobblesCmd.Flags().IntVarP(&limit, "limit", "l", 10, "Limit scrobble count")
	rootCmd.AddCommand(scrobblesCmd)
}
