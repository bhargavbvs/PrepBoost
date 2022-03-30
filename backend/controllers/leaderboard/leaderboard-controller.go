package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"prepboost.com/web/models"
)

func LeaderboardAlltime(w http.ResponseWriter, r *http.Request) {
	leadersAlltime := models.GetOverallLeaderboard()
	if err != nil {
		fmt.Println("error while parsing questions year")
	}

	body, err := json.Marshal(questions)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func LeaderboardDaily(w http.ResponseWriter, r *http.Request) {
	leaderboardDaily := models.GetQuestionsByYear(YEAR)
	if err != nil {
		fmt.Println("error while parsing questions year")
	}

	body, err := json.Marshal(questions)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
