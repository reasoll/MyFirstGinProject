package model

type AnswerMaster struct {
	BaseModel
	Answer     string
	QuestionId int
}

//根据question_id获取answer
func (this AnswerMaster) GetAnswerListByQuestionID(id int) (answerMasterList []AnswerMaster) {
	db.Where("question_id = ?", id).Find(&answerMasterList)
	return
}
