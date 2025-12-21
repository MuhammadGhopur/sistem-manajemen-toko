package config

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func LoadAppTemplates(r *gin.Engine) {
	var htmlFiles []string
	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			htmlFiles = append(htmlFiles, path)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Kesalahan saat memindai template: %v", err)
	}

	tmpl := template.New("").Funcs(r.FuncMap)
	tmpl, err = tmpl.ParseFiles(htmlFiles...)
	if err != nil {
		log.Fatalf("Kesalahan saat memparsing template: %v", err)
	}
	r.SetHTMLTemplate(tmpl)
}
