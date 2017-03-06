package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// // Movie Struct
// type Movie struct {
// 	Title  string `json:"title"`
// 	Rating string `json:"rating"`
// 	Year   string `json:"year"`
// }

// var movies = map[string]*Movie{
// 	"tt0076759": &Movie{Title: "Star Wars: A New Hope", Rating: "8.7", Year: "1977"},
// 	"tt0082971": &Movie{Title: "Indiana Jones: Raiders of the Lost Ark", Rating: "8.6", Year: "1981"},
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/movies", handleMovies).Methods("GET")
// 	router.HandleFunc("/movie/{imdbKey}", handleMovie).Methods("GET", "DELETE", "POST")
// 	http.ListenAndServe(":8080", router)
// }

// func handleMovie(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")
// 	vars := mux.Vars(req)
// 	imdbKey := vars["imdbKey"]

// 	switch req.Method {
// 	case "GET":
// 		movie, ok := movies[imdbKey]
// 		if !ok {
// 			res.WriteHeader(http.StatusNotFound)
// 			fmt.Fprint(res, string("Movie not found"))
// 		}
// 		outgoingJSON, error := json.Marshal(movie)
// 		if error != nil {
// 			log.Println(error.Error())
// 			http.Error(res, error.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		fmt.Fprint(res, string(outgoingJSON))
// 	case "DELETE":
// 		delete(movies, imdbKey)
// 		res.WriteHeader(http.StatusNoContent)
// 	case "POST":
// 		movie := new(Movie)
// 		decoder := json.NewDecoder(req.Body)
// 		error := decoder.Decode(&movie)
// 		if error != nil {
// 			log.Println(error.Error())
// 			http.Error(res, error.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		movies[imdbKey] = movie
// 		outgoingJSON, err := json.Marshal(movie)
// 		if err != nil {
// 			log.Println(error.Error())
// 			http.Error(res, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		res.WriteHeader(http.StatusCreated)
// 		fmt.Fprint(res, string(outgoingJSON))
// 	}
// }

// func handleMovies(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")
// 	outgoingJSON, error := json.Marshal(movies)
// 	if error != nil {
// 		log.Println(error.Error())
// 		http.Error(res, error.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	fmt.Fprint(res, string(outgoingJSON))
// }
