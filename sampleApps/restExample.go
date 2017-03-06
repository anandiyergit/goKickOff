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

// // Album Struct
// type Album struct {
// 	Title  string `json:"title"`
// 	Rating string `json:"rating"`
// 	Year   string `json:"year"`
// }

// var movies = map[string]*Movie{
// 	"tt0816692": &Movie{Title: "Interstellar", Rating: "8.7", Year: "2014"},
// 	"tt1375666": &Movie{Title: "Inception", Rating: "8.6", Year: "2010"},
// 	"tt0468569": &Movie{Title: "The dark Knight", Rating: "9.0", Year: "2008"},
// }

// var album = map[string]*Album{
// 	"tt0076778": &Album{Title: "1989:Taylor Swift", Rating: "8.7", Year: "2014"},
// 	"tt0082972": &Album{Title: "Cloud Nine:Kygo", Rating: "8.6", Year: "2016"},
// }

// func main() {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/movies", handleMovies).Methods("GET")
// 	router.HandleFunc("/album", handleAlbum).Methods("GET")
// 	router.HandleFunc("/movie/{imdbKey}", handleMovie).Methods("GET")
// 	router.HandleFunc("/album/{imdbKey}", handleAlb).Methods("GET")
// 	http.ListenAndServe(":8080", router)
// }

// func handleMovie(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")

// 	vars := mux.Vars(req)
// 	imdbKey := vars["imdbKey"]

// 	log.Println("Request for:", imdbKey)

// 	if movie, ok := movies[imdbKey]; ok {
// 		outgoingJSON, error := json.Marshal(movie)

// 		if error != nil {
// 			log.Println(error.Error())
// 			http.Error(res, error.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		fmt.Fprint(res, string(outgoingJSON))
// 	} else {
// 		res.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(res, string("Movie not found"))
// 	}
// }

// func handleAlb(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "applicatio/json")

// 	vars := mux.Vars(req)
// 	imdbKey := vars["imdbKey"]

// 	log.Println("Request for:", imdbKey)

// 	if alb, ok := album[imdbKey]; ok {
// 		outgoingJSON, error := json.Marshal(alb)

// 		if error != nil {
// 			log.Println(error.Error())
// 			http.Error(res, error.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		fmt.Fprint(res, string(outgoingJSON))
// 	} else {
// 		res.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(res, string("Album not found"))
// 	}
// }

// func handleMovies(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")

// 	outgoingJSON, error := json.Marshal(movies)
// 	// file, err := os.Create("sampleFile.json")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// file.Write(outgoingJSON)

// 	// file.Close()
// 	if error != nil {
// 		log.Println(error.Error())
// 		http.Error(res, error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprint(res, string(outgoingJSON))
// }

// func handleAlbum(res http.ResponseWriter, req *http.Request) {
// 	res.Header().Set("Content-Type", "application/json")

// 	outgoingJSON, error := json.Marshal(album)

// 	if error != nil {
// 		log.Println(error.Error())
// 		http.Error(res, error.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fmt.Fprint(res, string(outgoingJSON))
// }
