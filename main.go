package main

import (
	env "podcast/environment"
	"podcast/handler"
	//"podcast/middleware"
	"podcast/repository"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(env.GetLogLevel())
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/active", handler.StatusHandler).Methods(http.MethodGet, http.MethodOptions)

	router.HandleFunc("/podcasts", handler.AddPodcastHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/podcasts/{id}", handler.GetPodcastHandler).Methods(http.MethodGet)
	router.HandleFunc("/podcasts/{id}", handler.DeletePodcastHandler).Methods(http.MethodDelete)

	router.HandleFunc("/podcasts/{id}/episodes", handler.AddEpisodeHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/podcasts/{id}/episodes/{n}", handler.GetEpisodeHandler).Methods(http.MethodGet)
	router.HandleFunc("/podcasts/{id}/episodes/{n}", handler.DeleteEpisodeHandler).Methods(http.MethodDelete)


	//router.HandleFunc("/coverage/v1/postalCode/{id:[0-9]+}", handler.DeletePostalCodeHandler).Methods(http.MethodDelete, http.MethodOptions)

	router.Use(mux.CORSMethodMiddleware(router))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", env.Port), router))

	defer (*middleware.Repo).(repository.PostalCodeRepoImpl).Db.Close()
}