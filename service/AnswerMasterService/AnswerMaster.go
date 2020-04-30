package AnswerMasterService

import "MyFirstGinProject/model"

func GetAnswerListByQuestionID(id int) []model.AnswerMaster {
	return model.AnswerMaster{}.GetAnswerListByQuestionID(id)
}
