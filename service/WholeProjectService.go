package service

import (
	"MyFirstGinProject/service/AnswerMasterService"
	"MyFirstGinProject/service/ProjectMasterService"
	"MyFirstGinProject/service/QuestionMasterService"
	"MyFirstGinProject/view"
	"fmt"
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

func CreateSurvey(survey view.WholePriject) bool {
	project := survey.Project
	fmt.Print(project)

	for _, qCa := range survey.QuestionContainAnswerList {
		question := qCa.Question
		fmt.Print(question)
		for _, answer := range qCa.AnswerList {
			fmt.Print(answer)

		}

	}
	return true
}
