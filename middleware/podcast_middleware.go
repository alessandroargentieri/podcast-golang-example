package middleware

import (
    "podcast/model"
    "podcast/repository"
)

var Repo *repository.PodcastRepo

func init() {
	LoadConfigs()
}

func LoadConfigs() {

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(env.GetLogLevel())

	configs := make(map[string]string)
	configs["host"] = env.DbHost
	configs["port"] = fmt.Sprintf("%v", env.DbPort)
	configs["user"] = env.DbUser
	configs["password"] = env.DbPassword
	configs["dbname"] = env.DbName

	repo, err := repository.InitializeRepository(configs)
	if err != nil {
		log.Error(err.Error())
	}

	Repo = &repo

	log.Debug("Set Podcast repository. ")

}

func SetRepo(repo repository.PodcastRepo) {
	Repo = &repo
}
//~~~~~~~~

AddPodcast(podcast model.Podcast) (model.Podcast, error) {
	return (*Repo).Add(podcast)
}

GetAllPodcasts() ([]model.Podcast, error) {
	
}

GetPodcastById(id string) (model.Podcast, error) {}

DeletePodcastById(id string) error {}

AddEpisode(id string, episode model.Episode) (model.Episode, error) {}

GetAllEpisodesByPodcastId(id string) ([]model.Episode, error) {}

GetEpisodeByNumber(n int) (model.Episode, error) {}

DeleteEpisodeByPodcastIdAndEpisodeNumber(id string, n int) error {}
