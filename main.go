package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Resume struct {
	InputYaml string
	OutputFile string
	TemplateDir string
}

func check(e error){
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func defaultResume (resume Resume) Resume {
	if resume.InputYaml == "" {
		resume.InputYaml = "resume.yml"
	}
	if resume.OutputFile == "" {
		resume.OutputFile = "index.html"
	}
	if resume.TemplateDir == "" {
		resume.TemplateDir = "templates"
	}

	return resume
}

func BuildResume (r Resume) {
	resume := defaultResume(r)

	out, err := os.Create(resume.OutputFile)
	defer out.Close()
	check(err)

	data := map[string]interface{}{}

	file, err := ioutil.ReadFile(resume.InputYaml)
	check(err)

	err = yaml.Unmarshal(file, &data)
	check(err)

	t, err := template.ParseGlob(resume.TemplateDir + "/*")
	check(err)

	err = t.Execute(out, data)
	check(err)
}

func main() {
	resume := Resume{}
	BuildResume(resume)
}
