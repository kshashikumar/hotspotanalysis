package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	Commits    []Commithistory `json:"commithistory"`
	Allcommits float64         `json:"allcommits"`
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

func Getdata(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var re Recieved
	var se Sending
	decoder.Decode(&re)
	se.Authorname = re.Name
	se.Repositoryname = re.Repo
	d.Getjson(se.Authorname, se.Repositoryname)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&d); err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Getdata)
	http.ListenAndServe(":8090", r)
}
