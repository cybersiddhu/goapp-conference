package main

import (
	"github.com/cybersiddhu/gobioweb"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"webapps/conference/controller"
)

func main() {
	dir, _ := os.Getwd()
	t := &gobioweb.TemplateCache{
		GlobPattern: "*.tmpl",
		Path:        filepath.Join(dir, "templates"),
	}
	t.CacheEntries()
	s := sessions.NewCookieStore([]byte(getSecretId(15)))
	app := &gobioweb.App{Template: t.Cache, Session: s}

	logWriter := getLogWriter(dir)

	r := mux.NewRouter()
	r.NotFoundHandler = gobioweb.NewController(gobioweb.NotFound, app)
	r.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.Welcome, app),
	)).Methods("GET")

	//user and login handling routes
	u := r.PathPrefix("/users").Subrouter()
	u.Handle("/new", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.NewUser, app),
	)).Methods("GET")
	u.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.CreateUser, app),
	)).Methods("POST")

	r.Handle("/login", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.Login, app),
	)).Methods("GET")
	r.Handle("/login", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.CreateSession, app),
	)).Methods("POST")
	r.Handle("/logout", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.DeleteSession, app),
	)).Methods("DELETE")


	//abstract handling routes
	ar := r.PathPrefix("/abstracts").Subrouter()
	ar.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.ListAbstract, app),
	)).Methods("GET")

	ar.Handle("/", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.CreateAbstract, app),
	)).Methods("POST")

	ar.Handle("/new", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.NewAbstract, app),
	)).Methods("GET")

	ar.Handle("/{id:[0-9]+}", handlers.CombinedLoggingHandler(
		logWriter,
		gobioweb.NewController(controller.ShowAbstract, app),
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

func getSecretId(length int) string {
	//all alphabets
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	//make a slice out of it
	alphaSlice := strings.Split(alphabet, "")
	maxRange := len(alphabet)

	//slice that will hold the random string
	idx := make([]string, length)

	//Seed the random number generator
	rand.Seed(time.Now().Unix())

	//Now a random number within the alphabet range
	// and then pick one up and put it in slice
	for i := range idx {
		idx[i] = alphaSlice[rand.Intn(maxRange)]
	}
	return strings.Join(idx, "")
}
