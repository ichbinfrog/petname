package generator

import (
	"fmt"
	"testing"
)

func TestGeneratorNew(t *testing.T) {
	g := &Generator{}

	// Fail template creation
	g.New("{{.a}}", "name", '+')
	if len(g.TemplateInterface) != 0 {
		t.Errorf("Templating should fail when unknown variable {{.a}} is given")
	}

	// Error when template broken
	g.New("{{ .. }}", "name", '+')
	if len(g.TemplateInterface) != 0 {
		t.Errorf("Templating succeed when it should have failed {{ .. }}")
	}

	// Successfully create template
	g.New("{{ .Name }}{{ .Adjective }}{{ .Adjective }}", "name", '+')
	if len(g.TemplateInterface) == 0 {
		t.Errorf("Templating failed when it should succeed")
	}
}

func TestGeneratorGet(t *testing.T) {
	g := &Generator{}

	// Test successfull name generation
	g.New("{{ .Name }}{{ .Adjective }}{{ .Adjective }}", "name", '-')
	fmt.Println(g.Get())
}

func BenchmarkGeneratorGet(b *testing.B) {
	g := &Generator{}
	g.New("{{ .Name }}{{ .Adjective }}{{ .Adjective }}", "name", '-')

	for i := 0; i < b.N; i++ {
		fmt.Println(g.Get())
	}
}
