package cmd

import (
	petname "github.com/ichbinfrog/petname/pkg"
	"github.com/spf13/cobra"
)

var (
	port int
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Launch a petname server",
	Long:  `Launches a petname server at a given port at localhost`,
	Run: func(cmd *cobra.Command, args []string) {
		r := &petname.Router{}
		r.Start(port)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().IntVarP(&port, "port", "p", 8093, "port on which to launch")
}
