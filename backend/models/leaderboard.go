package models

func GetOverallLeaderboard() []TopicwiseQuestions {

	var leaderUsers []TopicwiseQuestions

	var leaderboardOverallQuery = ("SELECT users.id, users.username, users.email, users.mobile, " +
		"count(answered) as answered FROM user_answers JOIN users on " +
		"user_answers.user_id = users.id group by user_id ORDER BY answered DESC")

	DB.Raw(leaderboardOverallQuery).Find(&leaderUsers)
	DB.LogMode(true)
	return leaderUsers

}
