package generator

import (
	"fmt"
	"testing"
)

func TestGeneratorNew(t *testing.T) {
	g := &Generator{}

	// Fail template creation
	err := g.New("{{..}}", "name1")
	if err == nil {
		t.Errorf("[tpl {{..}}] Templating should have failed")
	}

	// Successfully create template
	err = g.New("{{ Adverb }}{{ Name }}{{ Adjective }}{{ Adjective }}", "name2")
	if err != nil {
		t.Errorf("[tpl {{ Adverb }}{{ Name }}{{ Adjective }}{{ Adjective }}] Templating failed with error %s", err.Error())
	}
}

func TestGeneratorGet(t *testing.T) {
	g := &Generator{}
	tpl := "{{ Adverb }}~{{ Name }}-{{ Adjective }}.{{ Adjective }}"
	// Test successful name generation
	g.New(tpl, "name")
	s, err := g.Get()
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Printf("[tpl %s] Got %s\n", tpl, s)
}

func BenchmarkGeneratorGet(b *testing.B) {
	g := &Generator{}
	tpl := "{{ Adverb }}--{{ Name }}&&{{ Adjective }}"
	g.New(tpl, "name")

	fmt.Printf("[tpl %s] Benchmarking...\n", tpl)
	for i := 0; i < b.N*b.N; i++ {
		_, err := g.Get()
		if err != nil {
			b.Errorf(err.Error())
		}
	}
}
