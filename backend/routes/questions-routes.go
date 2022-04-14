package routes

import (
	"github.com/gorilla/mux"
	controllers_que "prepboost.com/web/controllers/questions"
)

var RegisterQuestionsRoutes = func(router *mux.Router) {
	router.HandleFunc("/questions/{year}", controllers_que.GetYearwiseQuestions).Methods("GET")
	router.HandleFunc("/topicwise/{userId}", controllers_que.GetTopicwiseQuestions).Methods("GET")
	router.HandleFunc("/questions/report", controllers_que.ReportQuestion).Methods("POST")
}
