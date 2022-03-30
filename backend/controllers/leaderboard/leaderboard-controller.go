package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"prepboost.com/web/models"
)

func LeaderboardAlltime(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year := vars["year"]
	YEAR, err := strconv.ParseInt(year, 0, 0)
	questions := models.GetQuestionsByYear(YEAR)
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
	vars := mux.Vars(r)
	year := vars["year"]
	YEAR, err := strconv.ParseInt(year, 0, 0)
	questions := models.GetQuestionsByYear(YEAR)
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
