package main

import (
	"os"
	"fmt"
	"log"
	"errors"
	"strings"
	"io/ioutil"
	"text/template"
)

var (
	NoInputError = errors.New("Input length was 0")
)

var (
	ExitError = 1
	ExitIncorrectUsage = 64
)

func EnvMap() map[string]string {
	envMap := map[string]string{}

	for _, v := range os.Environ() {
		p := strings.SplitN(v, "=", 2)

		if len(p) != 2 {
			continue
		}

		envMap[p[0]] = p[1]
	}

	return envMap
}

func printSyntax() {
	fmt.Println("Syntax: echo <template> | envtpl [> <output_file>]")
}

func loadTemplate() (*template.Template, error) {
	tplBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	if len(tplBytes) == 0 {
		return nil, NoInputError
	}

	tpl, err := template.New("confTpl").Parse(string(tplBytes))
	if err != nil {
		return nil, err
	}

	return tpl, nil
}

func renderTemplate(tpl *template.Template, env map[string]string) error {
	if err := tpl.Execute(os.Stdout, env); err != nil {
		return err
	}

	return nil
}

func main() {
	tpl, err := loadTemplate()
	if err != nil {
		if err == NoInputError {
			printSyntax()
			os.Exit(ExitIncorrectUsage)
		} else {
			log.Panic(err)
			os.Exit(ExitError)
		}
	}

	if err := renderTemplate(tpl, EnvMap()); err != nil {
		log.Panic(err)
		os.Exit(ExitError)
	}
}
