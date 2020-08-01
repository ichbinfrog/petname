package cmd

import (
	"os"
	"strconv"

	"github.com/rs/zerolog/log"

	petname "github.com/ichbinfrog/petname/pkg"
	"github.com/spf13/cobra"
)

var (
	file, definition string
	retries          int
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		g, err := petname.NewGenerator(file, definition, retries)
		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Failed to initialize generator")
		}

		nameChan := make(chan *string)
		nb, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal().
				Str("nb", args[0]).
				Err(err).
				Msg("Invalid number given")
		}
		go g.Generate(nb, nameChan)
		for name := range nameChan {
			os.Stdout.WriteString(*name + "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.PersistentFlags().IntVarP(&retries, "retries", "r", 5, "amount of failed attempts before giving up")
	genCmd.PersistentFlags().StringVarP(&file, "seed", "s", "", "seed file for name, adjectives and adverbs")
	genCmd.PersistentFlags().StringVarP(&definition, "template", "t", "{{ Name }}-{{ Adverb }}", "go template for generating petname")
}
