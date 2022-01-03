package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/KarmaYume/prj_1/pkg/config"
	"github.com/KarmaYume/prj_1/pkg/handlers"
	"github.com/KarmaYume/prj_1/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNum = ":9090"

var app config.AppConfig
var session *scs.SessionManager

//main function
func main() {

	//change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("STARTING SERVER AT PORT NUMBER  %s", portNum))
	//_ = http.ListenAndServe(portNum, nil)

	srv := &http.Server{
		Addr:    portNum,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
