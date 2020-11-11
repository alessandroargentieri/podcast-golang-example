package handler

import (
	"podcast/dto"
	env "podcast/environment"
	"podcast/middleware"
	"podcast/model"
	"podcast/util"
	"encoding/json"
	"fmt"
	"errors"
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
	w.Write([]byte(`{"success": true, "uptime": "` + uptime().String() + `"}`))
}

func AddPodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    if r.Method == http.MethodOptions {
        return
    }

    var err error
    var podcast model.Podcast

    podcast, err = parseJsonPodcast(r.Body)
    if err!=nil {
        podcast, err = middleware.AddPodcast(podcast)
    } else {
    	err = errors.New("400 - Bad Request: error while parsing the json. ")
    }
    SendJsonResponse(w, []model.Podcast{ podcast }, err)
}

func GetPodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)

    var err error
    var podcasts []model.Podcast

    id, id_found := mux.Vars(r)["id"]
    if !id_found {
    	podcasts, err = middleware.GetAllPodcasts()
    }
    podcast, err = middleware.GetPodcastById(id)
    if err!=nil {
    	podcasts = append(podcasts, podcast)
    }

    SendJsonResponse(w, podcasts, err)
}
func DeletePodcastHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)

    var err error
    
    id, id_found := mux.Vars(r)["id"]
    if !id_found {
        err = errors.New("400 - Bad Request: Podcast id is missing. ")
    } else {
    	err = middleware.DeletePodcastById(id)
    }
    SendJsonResponse(w, []model.Podcast{ Id: id }, err)
}

func AddEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    if r.Method == http.MethodOptions {
        return
    }
    
    var err error
    var episode model.Episode
    
    id, id_found := mux.Vars(r)["id"]
    if !id_found {
        err = errors.New("400 - Bad Request: Podcast id is missing. ")
    } else {
        episode, err = parseJsonEpisode(r.Body)
        if err!=nil {
            episode, err = middleware.AddEpisode(id, episode)
        } else {
        	err = errors.New("400 - Bad Request: Error while parsing the json. ")
        }
    }
     SendJsonResponse(w, []model.Episode{ episode }, err)
}
func GetEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    
    var err error
    var episodes []model.Episode
    
    id, id_found := mux.Vars(r)["id"]
    if !id_found {
        err = errors.New("400 - Bad Request: Podcast id is missing. ")
    } else {
        n, n_found := mux.Vars(r)["n"]
        if !n_found {
       	    episodes, err = middleware.GetAllEpisodesByPodcastId(id)
        }
        episode, err := middleware.GetEpisodeByNumber(n)
        if err!=nil {
            episodes = append(episodes, episode)
        }
    }
    SendJsonResponse(w, episodes, err)
    
}
func DeleteEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Info(r)
    setHeaders(w)
    
    var err error

    id, id_found := mux.Vars(r)["id"]
    if !id_found {
        err = errors.New("400 - Bad Request: Podcast id is missing. ")
    } else {
        n, n_found := mux.Vars(r)["n"]
        if !n_found {
       	    err = errors.New("400 - Bad Request: Episode number is missing. ")
        }
        err = middleware.DeleteEpisodeByPodcastIdAndEpisodeNumber(id, n)
    }
    SendJsonResponse(w, []model.Episode{}, err)
}

func generateResponse(entities []interface{}, err error) {
	if entities==nil || err!=nil {
		return dto.ResponseWrapper{ Success: false, Error: err.Error()}
	}
	return dto.ResponseWrapper{ Success: true, Entities: entities }
}

func SendJsonResponse(w http.ResponseWriter, entities []interface{}, err error) {
	jsonBytes, _ := json.Marshal(generateResponse(entities, error))
	log.Debug(string(jsonBytes))
	if err!=nil {
		w.WriteHeader(getErrorCode(err))
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

func parseJsonPodcast(body io.Reader) (model.Podcast, error) {
	jsonDecoder := json.NewDecoder(body)

	now := time.Now()
	podcast := model.Podcast{
		Created: now,
		Updated: now,
	}
	err := jsonDecoder.Decode(&podcast)
	if err != nil {
		log.Error(err.Error())
		return podcast, errors.New("400 - Bad Request: " + err.Error())
	}
	return podcast, nil
}

func parseJsonEpisode(body io.Reader) (model.Episode, error) {
	jsonDecoder := json.NewDecoder(body)

	now := time.Now()
	episode := model.Episode{
		Created: now,
		Updated: now,
	}
	err := jsonDecoder.Decode(&episode)
	if err != nil {
		log.Error(err.Error())
		return episode, errors.New("400 - Bad Request: " + err.Error())
	}
	return postalCode, nil
}

func getErrorCode(err error) int {
	code, err := strconv.Atoi(err.Error()[0:3])
	if err!=nil {
		return 500
	}
	return code
}




