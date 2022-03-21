package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/trecknotrack/bookings/internal/config"
	"github.com/trecknotrack/bookings/internal/render"

	"github.com/trecknotrack/bookings/internal/handlers"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// //Home is the home page handler
// func Home(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "home.page.tmpl")
// }

// //About is the about page handler
// func About(w http.ResponseWriter, r *http.Request) {
// 	renderTemplate(w, "about.page.tmpl")
// }

// func renderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("Error parsing template: ", err)
// 		return
// 	}
// }

//main is the main application function
func main() {

	// change this to truw when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
