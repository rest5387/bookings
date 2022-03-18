package handlers

import (
	"encoding/json"
	"fmt"
	"log"
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

	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
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
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation renders the make-reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{})
}

// Generals renders the room page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "generals.page.tmpl", &models.TemplateData{})
}

// Majors renders the room page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "majors.page.tmpl", &models.TemplateData{})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("Start date is %s, End date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Panicln(err)
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the search availability page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send the data to the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{})
}
