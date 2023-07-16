package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// movies in-memory database
var moviesMap = make(map[string]Movie) // mapOf<String, Movie>

func getMovies(w http.ResponseWriter, req *http.Request) {
  var movies = []Movie {}
  responseBytes := new(bytes.Buffer)
  
  for movieId := range moviesMap {
    movies = append(movies, moviesMap[movieId])
  }
  json.NewEncoder(responseBytes).Encode(movies)
  w.Write(responseBytes.Bytes())
}

func createMovie(w http.ResponseWriter, req *http.Request) {
  var reqBody []byte
  reqBody, _ = ioutil.ReadAll(req.Body)
  req.Body.Close()

  var movie Movie
  if err := json.Unmarshal(reqBody, &movie); err != nil {
    fmt.Println("The err --> ",err)
    http.Error(w, "400 Bad Request", http.StatusBadRequest)
    return
  }
  
  moviesMap[movie.ID] = movie
  w.Write([]byte("Movie added successfully"))
}

func getMovie(w http.ResponseWriter, req *http.Request) {
  params := mux.Vars(req)
  id := params["id"]

  requestedMovie := moviesMap[id]

  responseBytes := new(bytes.Buffer)
  json.NewEncoder(responseBytes).Encode(requestedMovie)
  w.Write(responseBytes.Bytes())
}

func main() {
  router := mux.NewRouter()
  
  router.HandleFunc("/movies", createMovie).Methods("POST")
  router.HandleFunc("/movies", getMovies).Methods("GET")
  router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  
  fmt.Println("The server is listening on port 8080")
  if err := http.ListenAndServe("localhost:8080", router); err != nil {
    log.Fatal(err)
  }
}
