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
	topicId := vars["topicId"]
	subtopicId := vars["subtopicId"]
	tID, err := strconv.ParseInt(topicId, 0, 0)
	stID, err := strconv.ParseInt(subtopicId, 0, 0)
	questions := models.GetTopicwiseQuestions(tID, stID)
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

//Report question if it is incorrect
func ReportQuestion(w http.ResponseWriter, r *http.Request) {

	// vars := mux.Vars(r)
	// questionId := vars["questionId"]
	// qID, err := strconv.ParseInt(questionId, 0, 0)
	// // questions := models.GetTopicwiseQuestions(qID, )
	// if err != nil {
	// 	fmt.Println("error while parsing questions year")
	// }

	// body, err := json.Marshal(questions)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(body)
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
