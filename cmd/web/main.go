package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/theflyingdutch789/bookings/pkg/Render"
	"github.com/theflyingdutch789/bookings/pkg/config"
	"github.com/theflyingdutch789/bookings/pkg/handlers"
	"log"
	"net/http"
	"time"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false
	// Initialize a new session manager and configure the session lifetime.
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := Render.CreateTemplateCache()
	if err != nil {
		log.Fatal("template cache creation failed line 14 main")
	}

	app.TemplateCache = tc
	app.UseCache = false

	Render.NewTemplates(&app)
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	srv := http.Server{
		Addr:    ":3000",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Error serving template in serve and listen func line 33 main")
	}
}
