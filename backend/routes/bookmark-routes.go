package routes

import (
	"github.com/gorilla/mux"
	controllers_que "prepboost.com/web/controllers/questions"
)

var RegisterBookmarkRoutes = func(router *mux.Router) {
	router.HandleFunc("/bookmarks/{userId}", controllers_que.GetBookmarkedQuestions).Methods("GET")
}
