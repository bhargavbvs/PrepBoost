package models

type Questions struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Exam        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Year        uint   `gorm:"not null"`
	Question    string `gorm:"not null"`
	Answer      string `gorm:"not null"`
	Option1     string `gorm:"not null"`
	Option2     string `gorm:"not null"`
	Option3     string `gorm:"not null"`
	Option4     string `gorm:"not null"`
	Explanation string `gorm:"null"`
	Learning    string `gorm:"null"`
	Source      string `gorm:"null"`
}

type TopicwiseQuestions struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Exam        string `gorm:"not null"`
	Type        string `gorm:"not null"`
	Year        uint   `gorm:"not null"`
	Question    string `gorm:"not null"`
	Answer      string `gorm:"not null"`
	Answered    string `gorm:"not null"`
	Option1     string `gorm:"not null"`
	Option2     string `gorm:"not null"`
	Option3     string `gorm:"not null"`
	Option4     string `gorm:"not null"`
	Explanation string `gorm:"null"`
	Learning    string `gorm:"null"`
	Source      string `gorm:"null"`
}

func GetQuestionsByYear(Year int64) []Questions {
	var que []Questions

	//Query to fetch the questions of that specific year and exam
	var queQuery = ("SELECT questions.id, exams.exam, exams.type, questions.year, questions.question, " +
		"topics.topic, subtopics.subtopic, questions.answer, questions.option1, questions.option2," +
		"questions.option3, questions.option4, questions.explanation, questions.learning," +
		"questions.source FROM questions INNER JOIN topics ON questions.topic_id = topics.id " +
		"INNER JOIN subtopics ON questions.subtopic1_id = subtopics.id " +
		"INNER JOIN exams ON questions.exam_id = exams.id " +
		"WHERE exam_id = ?  and year = ? ORDER BY id")

	db.Raw(queQuery, 1, Year).Find(&que)
	db.LogMode(true)
	return que
}

//Get questions based on the specific topic that user wants
func GetTopicwiseQuestions(userId int64) []TopicwiseQuestions {
	var que []TopicwiseQuestions

	//Query to fetch the questions of that specific year and exam
	var topicQuery = ("SELECT questions.id, exams.exam, exams.type, questions.year, questions.question, " +
		"topics.topic, subtopics.subtopic, questions.answer, user_answers.answered, user_answers.correct, " +
		"questions.option1, questions.option2, questions.option3, questions.option4, questions.explanation, " +
		"questions.learning, questions.source FROM questions " +
		"INNER JOIN topics ON questions.topic_id = topics.id " +
		"INNER JOIN subtopics ON questions.subtopic1_id = subtopics.id " +
		"INNER JOIN exams ON questions.exam_id = exams.id " +
		"LEFT JOIN user_answers ON questions.id = user_answers.question_id " +
		"AND user_answers.user_id = ? WHERE exam_id = ? and questions.topic_id = ? and " +
		"questions.subtopic1_id = ? ORDER BY id")

	db.Raw(topicQuery, userId, 1, 1, 1).Find(&que)
	db.LogMode(true)
	return que
}

func GetBookmarkedQuestions(userId int64) []Questions {
	var Questions []Questions
	db.Where("Year=?", userId).Find(&Questions)
	return Questions
}
