package bootstrap

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig"
)

// Project contains the vars needed for bootstrapping
type Project struct {
	Image    string
	ImageTag string
	Module   string
}

// Execute to write / gen all files
func (p *Project) Execute(path string) error {
	err := write("bootstrap/src/.gitattributes", path)
	if err != nil {
		return err
	}

	err = write("bootstrap/src/.gitignore", path)
	if err != nil {
		return err
	}

	err = generate("bootstrap/src/README.md", path, "README.md", p)
	if err != nil {
		return err
	}

	err = generate("bootstrap/src/Makefile", path, "Makefile", p)
	if err != nil {
		return err
	}

	err = generate("bootstrap/src/Dockerfile", path, "Dockerfile", p)
	if err != nil {
		return err
	}

	err = write("bootstrap/src/resource.go", path)
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Join(path, "cmd", "check"), os.ModePerm)
	err = generate("bootstrap/src/cmd.go", path, filepath.Join("cmd", "check", "main.go"), struct {
		Module string
		Cmd    string
	}{Module: p.Module, Cmd: "Check"})
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Join(path, "cmd", "get"), os.ModePerm)
	err = generate("bootstrap/src/cmd.go", path, filepath.Join("cmd", "get", "main.go"), struct {
		Module string
		Cmd    string
	}{Module: p.Module, Cmd: "Get"})
	if err != nil {
		return err
	}

	os.MkdirAll(filepath.Join(path, "cmd", "put"), os.ModePerm)
	err = generate("bootstrap/src/cmd.go", path, filepath.Join("cmd", "put", "main.go"), struct {
		Module string
		Cmd    string
	}{Module: p.Module, Cmd: "Put"})
	if err != nil {
		return err
	}

	return nil
}

func write(src, dest string) error {
	d, err := content(src)
	if err != nil {
		return err
	}

	return file(filepath.Join(dest, path.Base(src)), d)
}

func generate(src, dest, name string, data interface{}) error {
	var output bytes.Buffer
	d, err := content(src)
	if err != nil {
		return err
	}

	var tpl = buildTemplate(d)
	if err := tpl.Execute(&output, data); err != nil {
		return err
	}

	return file(filepath.Join(dest, name), output.Bytes())
}

func content(src string) ([]byte, error) {
	a, err := Asset(src)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func buildTemplate(d []byte) *template.Template {
	return template.Must(template.New("base").Funcs(sprig.TxtFuncMap()).Parse(string(d)))
}

func file(f string, d []byte) error {
	err := ioutil.WriteFile(f, d, 0644)
	if err != nil {
		return err
	}

	return nil
}
