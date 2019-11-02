package server

import (
	"github.com/ichbinfrog/petname/pkg/generator"
)

// API represents an API access point
type API struct {
	Name      string              `json:"name"`
	Lock      bool                `json:"lock"`
	Token     []string            `json:"token,omitempty"`
	Generator generator.Generator `json:"generator,omitempty"`
}

// SetupAPI sets up an API
func (i *Instance) SetupAPI(name string, lock bool, template string, separator rune) bool {
	for _, a := range i.API {
		if name == a.Name {
			return false
		}
	}

	api := API{
		Name:      name,
		Lock:      lock,
		Token:     []string{},
		Generator: generator.Generator{},
	}
	api.Generator.New(template, name, separator)
	i.API = append(i.API, api)
	return true
}
