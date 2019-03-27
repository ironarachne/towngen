package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/random"
	"github.com/ironarachne/towngen"
)

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newTown towngen.Town

	random.SeedFromString(id)

	newTown = towngen.GenerateTown("random", "random")

	json.NewEncoder(w).Encode(newTown)
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var newTown towngen.Town

	rand.Seed(time.Now().UnixNano())

	newTown = towngen.GenerateTown("random", "random")

	json.NewEncoder(w).Encode(newTown)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", getTownRandom)
	r.Get("/{id}", getTown)

	fmt.Println("Town Generator API is online.")
	log.Fatal(http.ListenAndServe(":7979", r))
}
