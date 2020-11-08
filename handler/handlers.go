package handler

import (
	"podcast/dto"
	env "podcast/environment"
	"podcast/middleware"
	"podcast/model"
	"podcast/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	LoadConfigs()
}

func LoadConfigs() {
	startTime = time.Now()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(env.GetLogLevel())
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	setHeaders(w)
    if r.Method == http.MethodOptions {
        return
    }
	w.Write([]byte(`{"result": {"success": true, "uptime": "` + uptime().String() + `"}}`))
}

func AddPodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    if r.Method == http.MethodOptions {
        return
    }
}
func GetPodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)

    var podcasts []model.Podcast

    id, found := mux.Vars(r)["id"]
    if !found {
    	podcasts, err = middleware.GetAllPodcasts()
    }
    podcasts, err = middleware.GetPodcastsById(id)

    SendJsonResponse(w, podcasts, err)

}




func DeletePodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    pid, pid_found := mux.Vars(r)["id"]
}
func AddEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    if r.Method == http.MethodOptions {
        return
    }
}
func GetEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    pid, pid_found := mux.Vars(r)["pid"]
    eid, eid_found := mux.Vars(r)["eid"]
}
func DeleteEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    pid, pid_found := mux.Vars(r)["pid"]
    eid, eid_found := mux.Vars(r)["eid"]
}

func GenerateResponse(entities []interface{}, err error) {
	if entities==nil || err!=nil {
		return dto.ResponseWrapper{ Success: false, Error: err.Error()}
	}
	return dto.ResponseWrapper{ Success: true, Entities: entities }
}

func SendJsonResponse(w http.ResponseWriter, entities []interface{}, err error) {
	jsonBytes, _ := json.Marshal(GenerateResponse(entities, error))
	log.Debug(string(jsonBytes))
	if err!=nil {
		w.WriteHeader(500)
	}
  	w.Write([]byte(jsonBytes))
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-credentials", "true")
	w.Header().Set("access-control-allow-headers", "Authorization, Content-Type")
	w.Header().Set("access-control-expose-headers", "session")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}