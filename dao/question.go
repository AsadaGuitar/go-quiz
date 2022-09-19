package dao

import (
	"github.com/jinzhu/gorm"
)

type Questions struct {
	Id         int    `gorm:"column:id"`
	QuestionId int    `gorm:"column:question_id"`
	Text       string `gorm:"column:text"`
	Close      bool   `gorm:"column:close"`
}

func (Questions) TableName() string {
	return "questions"
}

type Answers struct {
	Id         int    `gorm:"column:id"`
	QuestionId int    `gorm:"column:question_id"`
	Text       string `gorm:"column:text"`
	Correct    bool   `gorm:"column:correct"`
	Close      bool   `gorm:"column:close"`
}

func (Answers) TableName() string {
	return "answers"
}

type Question struct {
	Id      int
	Text    string
	Answers []string
	Correct int
}

func NewQuestion(db *gorm.DB) *[]Question {

	var (
		questions []Question
		tables    []Questions
	)
	db.Where("close = ?", false).Find(&tables)

	for i := 0; i < len(tables); i++ {

		var answers []Answers
		db.Where("question_id = ?", tables[i].QuestionId).Find(&answers)

		q := Question{}
		q.Id = tables[i].Id
		q.Text = tables[i].Text

		var answerText []string
		for j := 0; j < len(answers); j++ {
			if answers[j].Correct {
				q.Correct = j
			}
			answerText = append(answerText, answers[j].Text)
		}
		q.Answers = answerText
		questions = append(questions, q)
	}

	return &questions
}
