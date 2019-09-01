package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger *loggo.Logger

type Env struct {
	config Config
}

func main() {
	var env Env
	env.config = CollectConfig()

	newLogger := loggo.GetLogger("main")
	logger = &newLogger

	err := loggo.ConfigureLoggers("<root>=TRACE")
	if err != nil {
		fmt.Printf("Error configurting Logger: %s", err.Error())
		return
	}

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(env.HandlePortal) // Top Router Catch All

	go http.ListenAndServe(":12345", r)

	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
	nch := make(chan os.Signal)
	signal.Notify(nch, syscall.SIGINT, syscall.SIGTERM)
	logger.Infof("%s", <-nch)
}
