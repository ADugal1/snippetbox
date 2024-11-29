package main

import (
	"html/template"
	"path/filepath"
)

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

		// Create slice containing the filepaths for base template, partials, and page
		files := []string{
			"./ui/html/base.tmpl",
			"./ui/html/partials/nav.tmpl",
			page,
		}

		// Parse files into a template set.
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map using the name of the page as key
		cache[name] = ts
	}

	// Return the map
	return cache, nil
}

// type templateData struct {
// 	Snippet  models.Snippet
// 	Snippets []models.Snippet
// }
