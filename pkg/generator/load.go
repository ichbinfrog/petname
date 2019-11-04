// Package generator encapsulates the structure which is in charge
// of populating the petname array
package generator

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// Adjectives is a list of default adjectives
	Adjectives = []string{}
	// Adverbs is a list of default adjectives
	Adverbs = []string{}
	// Names is a list of default names
	Names = []string{}
)

// Load setups up the configuration reload using viper
func Load(p string) error {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	// Search config in home directory with name ".petname" (without extension).
	viper.AddConfigPath(home)
	viper.AddConfigPath("/etc/petname")
	viper.AddConfigPath(p)
	viper.SetConfigName(".seed")

	if reloadErr := Reload(); err != nil {
		return reloadErr
	}
	return nil
}

// Reload makes viper reload the .seed configuration file
// that can be found either on $HOME/.seed, /etc/petname/.seed
// or ./.seed
func Reload() error {
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	name := viper.GetStringSlice("names")
	adv := viper.GetStringSlice("adverbs")
	adj := viper.GetStringSlice("adjectives")

	if len(name) != 0 {
		Names = name
	}

	if len(adv) != 0 {
		Adverbs = adv
	}

	if len(adj) != 0 {
		Adjectives = adj
	}
	return nil
}
