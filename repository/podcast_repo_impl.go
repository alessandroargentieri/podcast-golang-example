package repository

import (
  "context"
  "log"
  "time"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  env "podcast/environment"
)

func main() {
    connectionStr := fmt.Sprintf("mongodb://%s:%s/%s", env.DbHost, env.DbPort, env.DbName) //"mongodb://localhost:27017/mydbname"
    client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    quickstartDatabase := client.Database("quickstart")
    podcastsCollection := quickstartDatabase.Collection("podcasts")
    episodesCollection := quickstartDatabase.Collection("episodes")
}