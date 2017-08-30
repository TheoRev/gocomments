package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/TheoRev/gocomments/commons"
	"github.com/TheoRev/gocomments/migration"
	"github.com/TheoRev/gocomments/routes"
	"github.com/urfave/negroni"
)

// gocomments --migrate yes

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migración a la DB")
	flag.IntVar(&commons.Port, "port", 3030, "Puerto para el servidor web")
	flag.Parse()
	if migrate == "yes" {
		log.Println("Inició la migración...")
		migration.Migrate()
		log.Println("Finalizó la migración.")
	}

	router := routes.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", commons.Port),
		Handler: n,
	}

	log.Printf("Iniciado el servidor en http://localhost:%d", commons.Port)
	log.Println(server.ListenAndServe())
	log.Println("Finalizó la ejecución del programa")
}
