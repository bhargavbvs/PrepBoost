package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"prepboost.com/web/models"
)

func LeaderboardAlltime(w http.ResponseWriter, r *http.Request) {
	leadersAlltime := models.GetOverallLeaderboard()

	body, err := json.Marshal(leadersAlltime)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func LeaderboardDaily(w http.ResponseWriter, r *http.Request) {
	leaderboardDaily := models.GetDailyLeaderboard()

	body, err := json.Marshal(leaderboardDaily)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
