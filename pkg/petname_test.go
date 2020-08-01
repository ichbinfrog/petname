package petname

import (
	"testing"

	"github.com/rs/zerolog/log"
)

func TestTree(t *testing.T) {
	c := &Cache{}
	log.Print(c.Insert("ab"))

	log.Print(c)
	log.Print(c.Insert("ab"))
	log.Print(c)

}

func BenchmarkGenerator(b *testing.B) {
	g, err := NewGenerator("", "{{ Name }}----{{ Adverb }}", 5)
	if err != nil {
		log.Error().Err(err)
	}
	nameChan := make(chan *string)
	b.StartTimer()
	go g.Generate(10000, nameChan)
	for name := range nameChan {
		log.Info().Msg(*name)
	}
	b.StopTimer()
}

func TestServer(t *testing.T) {
	r := &Router{}
	r.Start(8093)
}
