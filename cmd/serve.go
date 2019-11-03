// Package cmd is the autogenerated root package by cobra
// to interface with the petname "cli"
/*
Copyright © 2019 ichbinfrog

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
	"fmt"
	"github.com/ichbinfrog/petname/internal/server"
	"github.com/spf13/cobra"
	"strconv"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launches a Petname API server",
	Long: `
Petname is a CLI application that can be used to generate a pronounceable, sometimes even memorable names consisting of a random combination of adverbs, an adjective, and an animal name.

This launches an API server which responds with a petname which runs on default at the /v1 endpoint. You can customize the endpoints by sending a GET with query values : name, lock, separator, template

Once this endpoint is setup, you can query:
- GET /get/{api} to get a single one
- GET /get/{api}/{n} to get n petnames in a json list`,
	Run: func(cmd *cobra.Command, args []string) {
		p, convErr := strconv.Atoi(port)
		if convErr != nil {
			panic(convErr)
		}

		if p <= 0 || p > 65535 {
			fmt.Errorf("Error launching API server, port must be within 0 and 65535")
		}

		i := &server.Instance{}
		fmt.Printf("Serving API petname server on port %d", p)
		i.Start(p)
	},
}

var (
	port string
)

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&port, "port", "p", port, "Port for serving the API server (default: 8000)")
}
