package routes

import (
	"github.com/gorilla/mux"
	controllers_que "prepboost.com/web/controllers/questions"
)

var RegisterQuestionsRoutes = func(router *mux.Router) {
	router.HandleFunc("/questions/{year}", controllers_que.GetYearwiseQuestions).Methods("GET")
}
