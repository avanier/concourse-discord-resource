// Copyright © 2017 Alexis Vanier <alexis@amonoid.io>
// This file is part of concourse-discord-resource.
//
// concourse-discord-resource is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or (at your
// option) any later version.
//
// concourse-discord-resource is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
// or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// concourse-discord-resource. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// outCmd represents the out command
var outCmd = &cobra.Command{
	Use:   "out",
	Short: "Send out a message",
	Long: `Out allows you to send a message to a channel or recipient.
It's the bread and butter of this resource.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("out called")
	},
}

func init() {
	RootCmd.AddCommand(outCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// outCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// outCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
