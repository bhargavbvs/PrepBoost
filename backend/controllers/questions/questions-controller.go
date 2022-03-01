package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"prepboost.com/web/models"
)

//Get questions based on the respective year
func GetYearwiseQuestions(w http.ResponseWriter, r *http.Request) {
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

//Get questions based on the respective topic
func GetTopicwiseQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	usID, err := strconv.ParseInt(userId, 0, 0)
	questions := models.GetTopicwiseQuestions(usID)
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

//Get questions bookmarked by the user according to the respective exam
func GetBookmarkedQuestions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	ID, err := strconv.ParseInt(userId, 0, 0)
	questions := models.GetBookmarkedQuestions(ID)
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
