package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.andrew.dugal/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map to act as a cache
	cache := map[string]*template.Template{}

	// Use the filepath.Gloc function to get a slice of all filepaths
	// that match the pattern below
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Loop through the page filepaths one by one
	for _, page := range pages {
		name := filepath.Base(page)

		// Parse the base template file into a template set
		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseGlob() *on this template set* to add the page template
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ParseFiles() *on this template set* to add the page template
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal
		cache[name] = ts
	}
	// Return the map
	return cache, nil
}
