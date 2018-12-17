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
	"github.com/spf13/cobra"
)

// searchCmd ...
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search the API",
}

// artistSearchCmd ...
var artistSearchCmd = &cobra.Command{
	Use:   "artist",
	Short: "Search the API for an artist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// albumSearchCmd ...
var albumSearchCmd = &cobra.Command{
	Use:   "album",
	Short: "Search the API for an album by artist and album name",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// TODO: Implement these commands
	searchCmd.AddCommand(artistSearchCmd)
	searchCmd.AddCommand(albumSearchCmd)
	rootCmd.AddCommand(searchCmd)
}
