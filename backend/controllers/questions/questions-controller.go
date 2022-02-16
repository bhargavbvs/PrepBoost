package controllers

// func GetYearwiseQuestions(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	year := vars["year"]
// 	var question []config.Questions
// 	result := db.Find(&question, "year = ?", year)
// 	body, err := json.Marshal(result)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(body)
// }
