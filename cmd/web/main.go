package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/DylanMorison/go-course/pkg/config"
	"github.com/DylanMorison/go-course/pkg/handlers"
	"github.com/DylanMorison/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting app on localhost%s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
