package repository

import (
  "context"
  "log"
  "time"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson/primitive"
  env "podcast/environment"
)

func init() {

}

const (
    DATABASE   = "podcast_db"
    COLLECTION = "podcasts"
)

type PodcastRepoImpl struct {
   Client *mongo.Client
  // Ctx    *context.timerCtx
}

func InitializeRepo() (PodcastRepo, error){
    connectionStr := fmt.Sprintf("mongodb://%s:%s", env.DbHost, env.DbPort) //"mongodb://localhost:27017/mydbname"
      //url := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, pass, host, port, db)
    client, err := mongo.NewClient(options.Client().ApplyURI(connectionStr))
    if err != nil {
        log.Error(err)
        return PodcastRepo{}, errors.New("503 - Service Unavailable: " + err.Error())
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Error(err)
        return PodcastRepo{}, errors.New("503 - Service Unavailable: " + err.Error())
    }
    var podcastRepo PodcastRepoImpl = PodcastCodeRepoImpl{ client }
    return podcastRepo, nil
   // defer client.Disconnect(ctx)
}


AddPodcast(repo PodcastRepoImpl) (podcast model.Podcast) (model.Podcast, error) {
    podcastsCollection := repo.Client.Database(DATABASE).Collection(COLLECTION)
    
    now := time.Now().UTC()

    podcast.Created = now
    podcast.Updated = now

    if len(podcast.Episodes)>0 {
        for i := 0; i < len(podcast.Episodes); i++ {
            podcast.Episodes[i].Created = now
            podcast.Episodes[i].Updated = now 
        }
    }
         
    result, err := podcastsCollection.InsertOne(ctx, podcast)
    newID := result.InsertedID
}

GetAllPodcasts() ([]model.Podcast, error) {

}

GetPodcastById(id string) (model.Podcast, error) {}

DeletePodcastById(id string) error {}

AddEpisode(id string, episode model.Episode) (model.Episode, error) {}

GetAllEpisodesByPodcastId(id string) ([]model.Episode, error) {}

GetEpisodeByNumber(n int) (model.Episode, error) {}

DeleteEpisodeByPodcastIdAndEpisodeNumber(id string, n int) error {}


func FindAll() {
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil { log.Fatal(err) }
    defer cur.Close(context.Background())
    for cur.Next(context.Background()) {
      // To decode into a struct, use cursor.Decode()
      result := struct{
        Foo string
        Bar int32
      }{}
      err := cur.Decode(&result)
      if err != nil { log.Fatal(err) }
      // do something with result...

      // To get the raw bson bytes use cursor.Current
      raw := cur.Current
      // do something with raw...
    }
    if err := cur.Err(); err != nil {
      return err
    }
}

func Find(filter bson.D{}) {
    result := struct{
      Foo string
      Bar int32
    }{}
    filter := bson.D{{"hello", "world"}}
    err := collection.FindOne(context.Background(), filter).Decode(&result)
    if err != nil { return err }
    // do something with result...
}

func DeleteOne() {
    var coll *mongo.Collection

    // delete at most one document in which the "name" field is "Bob" or "bob"
    // specify the SetCollation option to provide a collation that will ignore case for string comparisons
    opts := options.Delete().SetCollation(&options.Collation{
        Locale:    "en_US",
        Strength:  1,
        CaseLevel: false,
    })
    res, err := coll.DeleteOne(context.TODO(), bson.D{{"name", "bob"}}, opts)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("deleted %v documents\n", res.DeletedCount)
}

func FindOneAndUpdate() {
    var coll *mongo.Collection
    var id primitive.ObjectID

    // find the document for which the _id field matches id and set the email to "newemail@example.com"
    // specify the Upsert option to insert a new document if a document matching the filter isn't found
    opts := options.FindOneAndUpdate().SetUpsert(true)
    filter := bson.D{{"_id", id}}
    update := bson.D{{"$set", bson.D{{"email", "newemail@example.com"}}}}
    var updatedDocument bson.M
    err := coll.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
    if err != nil {
        // ErrNoDocuments means that the filter did not match any documents in the collection
        if err == mongo.ErrNoDocuments {
            return
        }
        log.Fatal(err)
    }
    fmt.Printf("updated document %v", updatedDocument)
}


func ReplaceOneOrInsert() {
    var coll *mongo.Collection
    var id primitive.ObjectID

    // find the document for which the _id field matches id and add a field called "location"
    // specify the Upsert option to insert a new document if a document matching the filter isn't found
    opts := options.Replace().SetUpsert(true)
    filter := bson.D{{"_id", id}}
    replacement := bson.D{{"location", "NYC"}}
    result, err := coll.ReplaceOne(context.TODO(), filter, replacement, opts)
    if err != nil {
        log.Fatal(err)
    }

    if result.MatchedCount != 0 {
        fmt.Println("matched and replaced an existing document")
        return
    }
    if result.UpsertedCount != 0 {
        fmt.Printf("inserted a new document with ID %v\n", result.UpsertedID)
    }
}


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




