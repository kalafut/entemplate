package entemplate_test

import (
	"html/template"
	"os"
	"strings"

	"github.com/kalafut/entemplate"
)

func ExampleExtemplate_ParseDir() {
	xt := entemplate.New().Funcs(template.FuncMap{
		"tolower": strings.ToLower,
	})
	_ = xt.ParseDir("examples/", []string{".tmpl"})
	_ = xt.ExecuteTemplate(os.Stdout, "child.tmpl", nil)
	// Output: Hello from child.tmpl
	//	Hello from partials/question.tmpl
}
