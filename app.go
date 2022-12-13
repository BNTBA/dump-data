package main

/*
	Author : Iordanis Paschalidis

	Date   : 01/15/2022
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Service object that contains the Port and Router of the application
type Service struct {
	Port   string
	Router *mux.Router
}

type Article struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Descpriotion string `json:"description"`
	Summary      string `json:"summary"`
	Img          string `json:"img"`
}

func GetLatestsArticles(w http.ResponseWriter, r *http.Request) {

	var articles []Article
	for i := 0; i < 10; i++ {
		article := Article{
			ID:           i,
			Title:        "Title " + string(i),
			Descpriotion: "Description " + string(i),
			Summary:      "Summary " + string(i),
			Img:          "https://www.hollywoodreporter.com/wp-content/uploads/2017/06/gettyimages-622492080_-_h_2017.jpg",
		}
		articles = append(articles, article)
	}
	jsonBody, _ := json.Marshal(articles)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(200)
	w.Write(jsonBody)

}

func GetArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	article := Article{
		ID:           1,
		Title:        "Title " + id,
		Descpriotion: "Description " + id,
		Summary:      "Summary " + id,
		Img:          "https://www.hollywoodreporter.com/wp-content/uploads/2017/06/gettyimages-622492080_-_h_2017.jpg",
	}

	jsonBody, _ := json.Marshal(article)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(200)
	w.Write(jsonBody)

}

/*
   Running the service in port 7673 (getting the value from ./assets/config/production.json )

       Endpoints:
		GET:
			/worldcup/general
		POST:
			"/club/season"
*/
func (s Service) run() {

	s.Port = ":7673"
	s.Router.HandleFunc(
		"/latest/articles",
		GetLatestsArticles,
	).Methods(http.MethodGet)

	s.Router.HandleFunc(
		"/article/{id}",
		GetArticle,
	).Methods(http.MethodGet)

	c := cors.New(cors.Options{})
	handler := c.Handler(s.Router)

	err := http.ListenAndServe(s.Port, handler)
	if err != nil {
		panic(fmt.Errorf("listener_and_serve %w", err))
	}
}

func main() {
	restService := Service{Router: mux.NewRouter().StrictSlash(true)}
	restService.run()
}
