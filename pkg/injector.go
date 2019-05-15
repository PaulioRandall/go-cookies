package pkg

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

// Injector represents a single template along with resources to populate it.
type Injector struct {
	Template  string // Path to the template file
	Resources string // Path to the root folder of injectable resources
}

// Inject is NOT designed to be called directly, instead it is used in templates
// and called by the "text/template" templating engine.
//
// It takes a filename that is relative to the Resources directory and returns
// its content with each line indented to the specified number of tabs. This is
// used within the Template to replace placeholders with some content.
func (i *Injector) Inject(filename string, indent int) string {
	path := i.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s := string(bytes)
	prefix := strings.Repeat("\t", indent)
	return i.forEachToken(s, "\n", func(i int, l string) string {
		return prefix + l
	})
}

// ForEachToken applies to each token within 's', that is delimited by 'sep',
// the function 'f'. The modified string is then returned.
//
// 'sep' and tokenisation behave exactly the as if calling
// 'strings.Split(s, sep)'.
func (i *Injector) forEachToken(s string, sep string, f func(i int, l string) string) string {
	tokens := strings.Split(s, sep)
	for i, l := range tokens {
		tokens[i] = f(i, l)
	}
	return strings.Join(tokens, sep)
}

// Compile creates the destination file by copying the template and filling the
// placeholders. Placeholders are relative references to files within the
// Template.Resources directory.
func (i *Injector) Compile(dst string) error {
	var err error

	t, err := template.ParseFiles(i.Template)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, i)
	if err != nil {
		os.Remove(f.Name())
		return err
	}

	return nil
}
