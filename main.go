package main

import (
	"github.com/cybersiddhu/gobioweb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"webapps/conference/controller"
	"github.com/gorilla/sessions"
)

func main() {
	dir, _ := os.Getwd()
	t := &gobioweb.TemplateCache{
		GlobPattern: "*.tmpl",
		Path:        filepath.Join(dir, "templates"),
	}
	t.CacheEntries()

	logWriter := getLogWriter(dir)

	r := mux.NewRouter()
	r.NotFoundHandler = gobioweb.NewController(gobioweb.NotFound, t.Cache)
	r.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.Welcome, t.Cache),
	)).Methods("GET")

	ar := r.PathPrefix("/abstracts").Subrouter()
	ar.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.ListAbstract, t.Cache),
	)).Methods("GET")

	ar.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.CreateAbstract, t.Cache),
	)).Methods("POST")

	ar.Handle("/new", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.NewAbstract, t.Cache),
	)).Methods("GET")

	ar.Handle("/{id:[0-9]+}", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.ShowAbstract, t.Cache),
	)).Methods("GET")

	http.Handle("/", r)
	url := "0.0.0.0:5000"
	log.Printf("running server on port %s", url)
	err := http.ListenAndServe(url, nil)
	if err != nil {
		log.Printf("error running server %v", err)
	}
}

func getLogWriter(d string) io.Writer {

	var logWriter io.Writer
	logPath := filepath.Join(d, "log")
	if _, err := os.Stat(logPath); err != nil {
		if os.IsNotExist(err) {
			logWriter = os.Stderr
			log.Print("sending to stderr")
		}
	} else {
		logFile := filepath.Join(logPath, "app.log")
		log.Printf("log file path is %s", logFile)

		if _, err := os.Stat(logFile); err != nil {
			if os.IsNotExist(err) {
				logWriter, err = os.Create(logFile)
				if err != nil {
					log.Fatal(err)
				}
				log.Print("opening a new app.log file")
			}
		} else {
		  //If log file exist write in append mode
			logWriter, err = os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND, 0664)
			if err != nil {
				log.Fatal(err)
			}
			log.Print("appending to a new app.log file")
		}
	}
	return logWriter
}
