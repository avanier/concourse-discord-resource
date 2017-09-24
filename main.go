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

package main

import (
	"log"

	"github.com/avanier/concourse-discord-resource/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
