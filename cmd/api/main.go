// Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// App Version = "1.0.0"
const version = "1.0.0"

type config struct {
	port int
	env  string // dev, stg, prd, etc...
}

// dependencies injections
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	//read in the flag that are needed to populate the config
	flag.IntVar(&cfg.port, "port", 4000, "API port")
	flag.StringVar(&cfg.env, "env", "dev", "dev, stg, prd")
	flag.Parse()

	//create a logger ~ use := for undeclared var
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	//create our http server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(), // point to our custom router in routes.go
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server at %s", cfg.env, srv.Addr)
	//start the server
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
