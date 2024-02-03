package main

import (
	"html/template"
	"path/filepath"

	"aaravdrive.com/snippetbox/pkg/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(filepath.Join(dir, "*.page.templ"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		fileName := filepath.Base(page)

		ts, err := template.ParseFiles(page)

		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.templ"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.templ"))
		if err != nil {
			return nil, err
		}

		cache[fileName] = ts

	}

	return cache, nil

}
