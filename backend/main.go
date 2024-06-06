package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recieved struct {
	Name string `json:"author"`
	Repo string `json:"reponame"`
}

type Sending struct {
	Authorname     string `json:"authorname"`
	Repositoryname string `json:"repositoryname"`
}

type Data struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Commits    []Commithistory    `json:"commithistory,omitempty" bson:"commithistory,omitempty"`
	Allcommits float64            `json:"allcommits,omitempty" bson:"allcommits,omitempty"`
	Reponame   string             `json:"reponame,omitempty" bson:"reponame,omitempty"`
	Authname   string             `json:"authorname,omitempty" bson:"authorname,omitempty"`
}

type Commithistory struct {
	Commit    float64 `json:"commit"`
	Author    string  `json:"author"`
	Additions float64 `json:"additions"`
	Deletions float64 `json:"deletions"`
	Changes   float64 `json:"changes"`
}

// func (s *Data) WailsInit(runtime *wails.Runtime) error {
// 	s.Getjson()
// 	return nil
// }
var d Data
var client *mongo.Client

// func (d *Data) Retdata() []Commithistory {

// 	return d.Commits
// }
// func (d *Data) RetCommits() float64 {
// 	return d.Allcommits
// }

func (d *Data) Getjson(author string, repo string) *Data {
	url := "https://api.github.com/repos/" + author + "/" + repo + "/stats/contributors"
	var results []map[string]interface{}
	var added, deleted, changed float64
	d.Reponame = repo
	d.Authname = author
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bytes, &results)
	resp.Body.Close()
	// fmt.Println(results)
	for _, val := range results {
		tot := val["total"].(float64)
		d.Allcommits = d.Allcommits + tot
		//fmt.Println(val["weeks"])
		week := val["weeks"].([]interface{})
		for _, val1 := range week {
			check := val1.(map[string]interface{})
			added = added + check["a"].(float64)
			deleted = deleted + check["d"].(float64)
		}
		changed = added + deleted
		auth := val["author"].(map[string]interface{})
		d.Commits = append(d.Commits, Commithistory{Commit: tot, Author: auth["login"].(string), Additions: added, Deletions: deleted, Changes: changed})
	}
	return &Data{}
}

func Getdatafromweb(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var re Recieved
	var se Sending
	decoder.Decode(&re)
	se.Authorname = re.Name
	se.Repositoryname = re.Repo
	d.Getjson(se.Authorname, se.Repositoryname)
	collection := client.Database("repoanalysis").Collection("wishlist")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	//collection.DeleteMany(ctx, bson.D{})
	result, _ := collection.InsertOne(ctx, d)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
	d = Data{}
}
func Getdata(w http.ResponseWriter, r *http.Request) {
	var da []Data
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	collection := client.Database("repoanalysis").Collection("wishlist")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var single Data
		cursor.Decode(&single)
		da = append(da, single)
		fmt.Println(da)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	json.NewEncoder(w).Encode(&da)
}

func Getdatabyid(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("im in id", id)
	var da Data
	collection := client.Database("repoanalysis").Collection("wishlist")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Data{ID: id}).Decode(&da)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	fmt.Println(da)
	json.NewEncoder(w).Encode(&da)
}
func Deletedatabyid(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("repoanalysis").Collection("wishlist")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err, _ := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error on deletion")
		return
	}
}
func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/", Getdatafromweb).Methods("POST")
	r.HandleFunc("/", Getdata).Methods("GET")
	r.HandleFunc("/commits/{id}", Getdatabyid).Methods("GET")
	r.HandleFunc("/{id}", Deletedatabyid).Methods("DELETE")
	http.ListenAndServe(":8090", handlers.CORS(headers, methods, origins)(r))
}
