package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbm"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Movies []Movie

func main() {
	r := mux.NewRouter()
	director := Director{
		FirstName: "Ivan",
		LastName:  "Amelin",
	}
	Movies = append(Movies, Movie{
		ID:       "1",
		Isbn:     "000000",
		Title:    "Puss in Boots",
		Director: &director,
	})
	Movies = append(Movies, Movie{
		ID:    "2",
		Isbn:  "000001",
		Title: "Men in Black",
		Director: &Director{
			FirstName: "Arina",
			LastName:  "Artsiomava",
		},
	})
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieByID).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovieByID).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovieByID).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func deleteMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, movie := range Movies {
		if movie.ID == params["id"] {
			Movies = append(Movies[:i], Movies[i+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(Movies)
	if err != nil {
		return
	}

}

func updateMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range Movies {
		if item.ID == params["id"] {
			Movies = append(Movies[:i], Movies[i+1:]...)
			var movie Movie
			err := json.NewDecoder(r.Body).Decode(&movie)
			if err != nil {
				return
			}
			movie.ID = params["id"]
			Movies = append(Movies, movie)
			err1 := json.NewEncoder(w).Encode(movie)
			if err1 != nil {
				return
			}
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		return
	}
	movie.ID = strconv.Itoa(rand.Intn(999999))
	Movies = append(Movies, movie)
	err1 := json.NewEncoder(w).Encode(movie)
	if err1 != nil {
		return
	}
}

func getMovieByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, movie := range Movies {
		if movie.ID == params["id"] {
			err := json.NewEncoder(w).Encode(movie)
			if err != nil {
				return
			}
		}
	}
}

func getAllMovies(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Movies)
	if err != nil {
		return
	}
}
