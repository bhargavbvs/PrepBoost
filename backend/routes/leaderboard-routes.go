package routes

import (
	"github.com/gorilla/mux"
	controllers "prepboost.com/web/controllers/leaderboard"
)

var RegisterLeaderboardRoutes = func(router *mux.Router) {
	router.HandleFunc("/leaderboard/alltime/", controllers.LeaderboardAlltime).Methods("POST")
	//router.HandleFunc("/leaderboard/alltimetopten/", controllers.GetOverallLeaderboardtopten).Methods("POST")
	router.HandleFunc("/leaderboard/daily/", controllers.LeaderboardDaily).Methods("POST")
	//router.HandleFunc("/leaderboard/dailytopten", controllers.LeaderboardDailytopten).Methods("POST")
}
