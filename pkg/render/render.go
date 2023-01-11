package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gauravdjga/booking_app/pkg/config"
	"github.com/gauravdjga/booking_app/pkg/models"
)

// func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

// 	if err != nil {
// 		log.Println("Error in rendering template", err)
// 	}

// 	err = parsedTemplate.Execute(w, nil)

// 	if err != nil {
// 		log.Println("Error in Parsing template", err)
// 		return
// 	}

// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]

// 	if !inMap {
// 		log.Println("Creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		log.Println("Using Cached Template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, tmpl)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{"./templates/" + t, "./templates/base.layout.tmpl"}

// 	tmpl, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil

// }

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func addDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var ts map[string]*template.Template
	var err error

	if app.UseCache {
		ts = app.TemplateCache
	} else {
		ts, err = CreateTemplateCache()
		if err != nil {
			log.Fatal(err)
		}
	}

	t, ok := ts[tmpl]
	if !ok {
		log.Fatal("Error getting the template cache")
	}

	buff := new(bytes.Buffer)
	td = addDefaultData(td)
	err = t.Execute(buff, td)
	if err != nil {
		log.Fatal(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	return myCache, err
}
