package handler

import (
	"net/http"
	// "fmt"
	"github.com/gauravdjga/booking_app/pkg/config"
	"github.com/gauravdjga/booking_app/pkg/models"
	"github.com/gauravdjga/booking_app/pkg/render"
)

var Repo *Repostitory

type Repostitory struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repostitory {
	return &Repostitory{
		App: a,
	}
}

func NewHandlers(r *Repostitory) {
	Repo = r
}



func (m *Repostitory) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}


func (m *Repostitory) About(w http.ResponseWriter, r *http.Request) {

	stringMap := map[string]string{"test":"hello world"}

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
