package config

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoadAppTemplates(r *gin.Engine) {
	tmpl := template.New("").Funcs(r.FuncMap)

	err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			name := strings.TrimPrefix(path, "templates") 
			name = strings.ReplaceAll(name, "\\", "/")       
			if strings.HasPrefix(name, "/") || strings.HasPrefix(name, "\\") { 
				name = name[1:]
			}

			_, err = tmpl.New(name).Parse(string(content))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Kesalahan saat memindai dan memparsing template: %v", err)
	}

	r.SetHTMLTemplate(tmpl)
}
