// Package generator encapsulates the structure which is in charge
// of populating the petname array
package generator

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/ichbinfrog/petname/pkg/dict"
)

// Generator encapsulates all functions that allow for the unique
// generation of a petname
type Generator struct {
	Name              string
	Used              *dict.Tree
	TemplateInterface []func() (int, string)
	Separator         rune
}

const (
	// NameTemplate is the call value for templating names
	NameTemplate = ".Name"
	// AdjectiveTemplate is the call value for templating adjectives
	AdjectiveTemplate = ".Adjective"
	// AdverbTemplate is the call value for templating adverbs
	AdverbTemplate = ".Adverb"
)

var (
	templateReg = regexp.MustCompile(`\{\{\s*(\.[a-zA-Z]*)\s*\}\}`)
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func getName() (int, string) {
	index := rand.Intn(len(Names))
	return index, Names[index]
}

func getAdjective() (int, string) {
	index := rand.Intn(len(Adjectives))
	return index, Adjectives[index]
}

func getAdverb() (int, string) {
	index := rand.Intn(len(Adverbs))
	return index, Adverbs[index]
}

// New generates a new generator from a given template
// If template does not follow the const given will return empty generator
func (g *Generator) New(t string, n string, s rune) {
	for _, v := range templateReg.FindAllStringSubmatch(t, -1) {
		if v[1] == NameTemplate {
			g.TemplateInterface = append(g.TemplateInterface, getName)
		} else if v[1] == AdjectiveTemplate {
			g.TemplateInterface = append(g.TemplateInterface, getAdjective)
		} else if v[1] == AdverbTemplate {
			g.TemplateInterface = append(g.TemplateInterface, getAdverb)
		} else {
			return
		}
	}

	if _, err := template.New(n).Parse(t); err != nil {
		fmt.Println(err)
		return
	}
	g.Separator = s
	g.Used = &dict.Tree{}
}

// Get generates an unique petname and returns that string
func (g *Generator) Get() string {
	index := make([]int, len(g.TemplateInterface))
	name := make([]string, len(g.TemplateInterface))
	for {
		for i, t := range g.TemplateInterface {
			index[i], name[i] = t()
		}

		if !g.Used.Search(index) {
			g.Used.Insert(index)
			goto success
		}
	}

success:
	return fmt.Sprintf(strings.Join(name, string(g.Separator)))
}
