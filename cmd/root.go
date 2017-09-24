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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/themalkolm/venom"
	yaml "gopkg.in/yaml.v2"
)

var RootCmd = &cobra.Command{
	Use:          "example",
	Short:        "Do example things.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var cfg Config
		err := venom.Unmarshal(&cfg, viper.GetViper())
		if err != nil {
			return err
		}
		return runE(&cfg)
	},
}

type Config struct {
	ClientID     string
	ClientSecret string
	BotUsername  string
	BotToken     string
}

func init() {
	RootCmd.PersistentFlags().String("client-id", "", "A Discord App Client ID must be set")
	RootCmd.PersistentFlags().String("client-secret", "", "A Discord App Client Secret must be set")
	RootCmd.PersistentFlags().String("bot-username", "", "A Discord Bot Username must be set")
	RootCmd.PersistentFlags().String("bot-token", "", "A Discord Bot Token must be set")

	defaults := Config{
		ClientID:     "",
		ClientSecret: "",
		BotUsername:  "",
		BotToken:     "",
	}

	flags := venom.MustDefineFlags(defaults)
	RootCmd.PersistentFlags().AddFlagSet(flags)

	err := venom.TwelveFactorCmd("example", RootCmd, RootCmd.PersistentFlags())
	if err != nil {
		log.Fatal(err)
	}
}

func runE(cfg *Config) error {
	b, err := yaml.Marshal(&cfg)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", string(b))

	return nil
}
