package handlers

import (
	"github.com/theflyingdutch789/bookings/models"
	"github.com/theflyingdutch789/bookings/pkg/Render"
	"github.com/theflyingdutch789/bookings/pkg/config"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	stringMap := map[string]string{}
	stringMap["test"] = "Nabeel is a awesome guy"
	Render.RenderTemplate(res, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	//renderTemplate(res, "about.page.tmpl")
	stringMap := map[string]string{}
	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	Render.RenderTemplate(res, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
