package petname

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"text/template"

	"encoding/json"

	"github.com/gobuffalo/packr"
	"github.com/rs/zerolog/log"
)

type Cache map[string]struct{}

func (c Cache) Insert(name string) bool {
	if _, ok := c[name]; ok {
		return false
	}
	c[name] = struct{}{}
	return true
}

func (c Cache) Free(name string) {
	delete(c, name)
}

func (c Cache) Clear() {
	c = make(map[string]struct{})
}

type Generator struct {
	Adj  []string `json:"adjectives"`
	Adv  []string `json:"adverbs"`
	Name []string `json:"names"`

	Template *template.Template `json:"template"`
	Cache    *Cache             `json:"cache"`
	Retries  int
}

func (g *Generator) getName() string {
	return g.Name[rand.Intn(len(g.Name))]
}

func (g *Generator) getAdjective() string {
	return g.Adj[rand.Intn(len(g.Adj))]
}

func (g *Generator) getAdverb() string {
	return g.Adv[rand.Intn(len(g.Adv))]
}

func NewGenerator(file, definition string, retries int) (*Generator, error) {
	g := &Generator{
		Cache:   &Cache{},
		Retries: retries,
	}

	var err error
	g.Template, err = template.New("petname").Funcs(template.FuncMap{
		"Name":      func() string { return g.getName() },
		"Adverb":    func() string { return g.getAdjective() },
		"Adjective": func() string { return g.getAdverb() },
	}).Parse(definition)

	var data []byte
	if file == "" {
		box := packr.NewBox("../resources/")
		data, err = box.Find(".seed.json")
	} else {
		data, err = ioutil.ReadFile(file)
	}
	if err != nil {
		log.Error().
			Err(err).
			Str("file", file).
			Msg("Failed to read seed")
		return nil, err
	}

	if err := json.Unmarshal(data, g); err != nil {
		log.Error().
			Str("file", file).
			Msg("Failed unmarshal json")
		return nil, err
	}
	return g, nil
}

func (g *Generator) Generate(n int, name chan *string) {
	buf := &bytes.Buffer{}
	for i := 0; i < n; i++ {
		retries := g.Retries
		for {
			buf.Reset()
			if retries <= 0 {
				break
			}
			if err := g.Template.Execute(buf, nil); err != nil {
				break
			}
			petname := buf.String()
			if g.Cache.Insert(petname) {
				name <- &petname
				break
			}
			retries--
		}
		buf.Reset()
	}
	close(name)
}
