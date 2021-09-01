package handlers

import (
	"fmt"
	"net/http"

	"github.com/trecknotrack/bookings/pkg/models"

	"github.com/trecknotrack/bookings/pkg/config"

	"github.com/trecknotrack/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	app *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		app: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.app.Session.Put(r.Context(), "remote_ip", remoteIP)
	fmt.Println("Home put remote addr", m.app.Session.GetString(r.Context(), "remote_ip"))
	fmt.Println("remote_ip", r.RemoteAddr)
	fmt.Println("(Home func)Session key remote_ip exists?", m.app.Session.Exists(r.Context(), "remote_ip"))

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."
	fmt.Println("Session lifetime", m.app.Session.Lifetime)
	fmt.Println("Session cookie name", m.app.Session.Cookie.Name)
	fmt.Println("Session key remote_ip exists?", m.app.Session.Exists(r.Context(), "remote_ip"))
	remoteIP := m.app.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	fmt.Println("Actual remote addr", r.RemoteAddr)
	fmt.Println("About get remote addr", remoteIP, "\nstringmap ", stringMap["remote_ip"])

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
