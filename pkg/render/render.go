package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate is a method to render the HTML template for the application
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Cache template
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache:", err)
	}

	// Get template from cache
	temp, ok := tc[tmpl]
	if !ok {
		log.Fatal("Error loading template:", tmpl)
	}

	// buffer for more error checking
	buffer := new(bytes.Buffer)
	err = temp.Execute(buffer, data)
	if err != nil {
		log.Println("Error executing template:", err)
	}

	// Render Template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println("Error executing template:", err)
	}
}

// createTemplateCache is used for creating a template cache using a map and returns the map and error
func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all the files of the pages in templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// Iterate through all the pages ending with .page.html in the pages slice
	for _, page := range pages {
		// name stored the last element of path that is returned by using the filepath.Base.
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		// check for layout file in the templates folder
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
