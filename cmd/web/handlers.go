package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.andrew.dugal/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Call the newTemplateData() helper to get a templateData struct containing
	// the 'default' data, and add the snippets slice to it
	data := app.newTemplateData(r)
	data.Snippets = snippets

	// Pass the data to render() helper as normal
	app.render(w, r, http.StatusOK, "home.tmpl", data)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Same as above in the home handler
	data := app.newTemplateData(r)
	data.Snippet = snippet

	// Pass the data to render() helper as normal
	app.render(w, r, http.StatusOK, "view.tmpl", data)
}

// 	// Initialize a slice containing the paths to the view.tmpl file,
// 	// plus the base layout and navigation partial that we made earlier.
// 	files := []string{
// 		"./ui/html/base.tmpl",
// 		"./ui/html/partials/nav.tmpl",
// 		"./ui/html/pages/view.tmpl",
// 	}

// 	ts, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 		return
// 	}

// 	data := templateData{
// 		Snippet: snippet,
// 	}

// 	err = ts.ExecuteTemplate(w, "base", data)
// 	if err != nil {
// 		app.serverError(w, r, err)
// 	}

// 	// fmt.Fprintf(w, "%+v", snippet)
// }

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// dummy values to be removed later
	title := "0 snail"
	content := "0 snail\nClimbn Mount Fuji,\nBut slowl, slowly!\n\n- Kobayashi Issa"
	expires := 7

	// Pass the data to SnippetModel.Insert() method
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	// Return user to the relevant page for the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	// Not sure if these are staying?
	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte("Save a new snippet..."))
}
