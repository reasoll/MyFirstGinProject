package QuestionMasterService

import "MyFirstGinProject/model"

func GetQuestionMasterList() []model.QuestionMaster {

	return model.QuestionMaster{}.GetQuestionMasterAll()

}

func GetQuestionMasterById(id int) []model.QuestionMaster {

	return model.QuestionMaster{}.GetQuestionMasterById(id)
}
