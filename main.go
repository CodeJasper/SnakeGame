package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"snakeGameApi/config"
	"snakeGameApi/models"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	config.ConfigDB()
	defer config.Connect.Close(context.Background())

	r := chi.NewRouter()

	r.Route("/scores", func(r chi.Router) {
		r.Get("/", getScores)
		r.Post("/", addScore)
	})

	http.ListenAndServe(":8080", r)
	fmt.Println("Estado del servidor")
}

func getScores(w http.ResponseWriter, r *http.Request) {
	var scores []models.Score
	response, err := config.Connect.Query(context.Background(), "select * from scores order by score desc limit 10;")
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		for response.Next() {
			var scoreRow models.Score
			var id, score int
			var username string
			if err := response.Scan(&id, &score, &username); err != nil {
				var error models.Error
				error.Error = err.Error()
				json.NewEncoder(w).Encode(error)
				log.Println(err)
			} else {
				scoreRow.Score = score
				scoreRow.Username = username
				scores = append(scores, scoreRow)
			}
		}
		json.NewEncoder(w).Encode(scores)
	} else {
		var error models.Error
		error.Error = err.Error()
		json.NewEncoder(w).Encode(error)
		log.Println(err)
	}

}

func addScore(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newScore models.Score
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		var error models.Error
		error.Error = err.Error()
		json.NewEncoder(w).Encode(error)
		log.Println(err)
	} else {
		json.Unmarshal(reqBody, &newScore)
		response, err := config.Connect.Query(context.Background(), "INSERT INTO SCORES(score, username) VALUES("+strconv.Itoa(newScore.Score)+",'"+newScore.Username+"');")
		if err != nil {
			var error models.Error
			error.Error = err.Error()
			json.NewEncoder(w).Encode(error)
			log.Println(err)
		} else {
			defer response.Close()
			var success models.Successful
			success.Message = "Se ha registrado satisfactoriamente el score"
			json.NewEncoder(w).Encode(success)
			log.Println(success)
		}
	}
}
