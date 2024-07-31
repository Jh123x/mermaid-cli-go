package handler

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/jh123x/mermaid-cli-go/internal/common"
)

func GenHTML(tmplCfg *common.Template, filePath string) error {
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}

	tmplPath := path.Join(currDir, common.TmplDir, common.TmplFile)
	tmpl, err := template.New(path.Base(tmplPath)).ParseFiles(tmplPath)
	if err != nil {
		return err
	}

	fmt.Println("tmp path:", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, tmplCfg); err != nil {
		return err
	}

	return nil
}
