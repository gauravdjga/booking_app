package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	handler "github.com/gauravdjga/booking_app/pkg/handlers"
	"github.com/gauravdjga/booking_app/pkg/render"
	"github.com/gauravdjga/booking_app/pkg/config"
)

const portNumber = ":8080"

var session *scs.SessionManager

func main() {

	var app config.AppConfig

	app.InProduction = false 

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = true
	session.Cookie.Secure = app.InProduction

	app.Session = session

	ts, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = ts
    render.NewTemplates(&app)
	app.UseCache = false

	Repo := handler.NewRepo(&app)
	handler.NewHandlers(Repo)

	fmt.Println("Starting Application on port ", portNumber)

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
