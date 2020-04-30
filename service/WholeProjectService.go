package service

import (
	"MyFirstGinProject/service/AnswerMasterService"
	"MyFirstGinProject/service/ProjectMasterService"
	"MyFirstGinProject/service/QuestionMasterService"
	"MyFirstGinProject/view"
)

func GetWholeProject(id int) (project view.WholePriject) {

	project.Project = ProjectMasterService.GetProjectMaster(id)

	qmList := QuestionMasterService.GetQuestionMasterById(id)

	var qCa = view.QuestionContainAnswer{}
	for _, qm := range qmList {

		amList := AnswerMasterService.GetAnswerListByQuestionID(qm.ID)
		qCa.Question = qm
		qCa.AnswerList = amList
		project.QuestionContainAnswerList = append(project.QuestionContainAnswerList, qCa)

	}

	return project
}
