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

// Inject takes a filename that is relative to the Resources directory and
// returns its content with each line indented to the specified number of tabs.
// This is used within the Template to replace placeholders with some content.
func (i *Injector) Inject(filename string, indent int) string {
	path := i.Resources + filename
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	s := string(bytes)
	lines := strings.Split(s, "\n")
	prefix := strings.Repeat("\t", indent)

	for i, l := range lines {
		lines[i] = prefix + l
	}

	r := strings.Join(lines, "\n")
	return r
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
