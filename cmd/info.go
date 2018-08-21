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

// infoCmd represents the info command
var artistInfoCmd = &cobra.Command{
	Use:   "artist",
	Short: "Get info about an artist",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		artistInfo, err := client.GetArtistInfo(query)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(artistInfo)
	},
}

// infoCmd represents the info command
var albumInfoCmd = &cobra.Command{
	Use:   "album",
	Short: "Get info about an album",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		artist, album := args[0], args[1]
		albumInfo, err := client.GetAlbumInfo(artist, album)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(albumInfo)
	},
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about an artist or an album",
}

func init() {
	infoCmd.AddCommand(artistInfoCmd)
	infoCmd.AddCommand(albumInfoCmd)
	rootCmd.AddCommand(infoCmd)
}
