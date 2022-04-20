package models

type Leaderboard struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null"`
	Mobile   string `gorm:"not null"`
	Answered string `gorm:"not null"`
}

func GetOverallLeaderboard() []Leaderboard {

	var leaderUsers []Leaderboard

	var leaderboardOverallQuery = ("SELECT users.id, users.username, users.email, users.mobile, " +
		"count(answered) as answered FROM user_answers JOIN users on " +
		"user_answers.user_id = users.id group by user_id ORDER BY answered DESC")

	DB.Raw(leaderboardOverallQuery).Find(&leaderUsers)
	DB.LogMode(true)
	return leaderUsers
}

func GetOverallLeaderboardtopten() []Leaderboard {

	var leaderUsers []Leaderboard

	var leaderboardOverallQuery = ("SELECT users.id, users.username, users.email, users.mobile, " +
		"count(answered) as answered FROM user_answers JOIN users on " +
		"user_answers.user_id = users.id group by user_id ORDER BY answered DESC LIMIT 10")

	DB.Raw(leaderboardOverallQuery).Find(&leaderUsers)
	DB.LogMode(true)
	return leaderUsers
}



func GetDailyLeaderboard() []Leaderboard {

	var leaderDailyUsers []Leaderboard

	var leaderboardOverallQuery = ("SELECT users.id, users.username, users.email, users.mobile, " +
		"count(answered) as answered FROM user_answers JOIN users on " +
		"user_answers.user_id = users.id WHERE prepboost.user_answers.created_at > SUBDATE(NOW(), 1) " +
		"group by user_id ORDER BY answered DESC")

	DB.Raw(leaderboardOverallQuery).Find(&leaderDailyUsers)
	DB.LogMode(true)
	return leaderDailyUsers
}
