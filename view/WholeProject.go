package view

import "MyFirstGinProject/model"

type WholePriject struct {
	Project                   model.ProjectMaster
	QuestionContainAnswerList []QuestionContainAnswer
}

type QuestionContainAnswer struct {
	Question   model.QuestionMaster
	AnswerList []model.AnswerMaster
}
