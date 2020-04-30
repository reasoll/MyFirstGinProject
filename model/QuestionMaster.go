package model

type QuestionMaster struct {
	BaseModel
	Question        string
	CorrentAnswerId int
	ProjectId       int
}

//获取所有问题
func (this QuestionMaster) GetQuestionMasterAll() (questionMasterList []QuestionMaster) {

	db.Find(&questionMasterList)
	return
}

/**
根据Project_id获取Question
*/
func (this QuestionMaster) GetQuestionMasterById(id int) (questionMasterList []QuestionMaster) {

	db.Where("project_id = ?", id).Find(&questionMasterList)
	return
}
