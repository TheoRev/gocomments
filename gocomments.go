package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/TheoRev/DemoGoApp/migration"
	"github.com/TheoRev/gocomments/routes"
	"github.com/urfave/negroni"
)

// gocomments --migrate yes

func main() {
	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migraci+on a la DB")
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
		Addr:    "3030",
		Handler: n,
	}

	log.Println("Iniciado el servidor en http://localhost:3030")
	log.Println(server.ListendAndServe())
	log.Println("Finalizó la ejecución del programa")
}
