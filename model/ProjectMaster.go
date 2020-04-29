package model

type ProjectMaster struct {
	BaseModel
	ProjectName		string
}


//Create
func (this *ProjectMaster) Add() {
	db.Create(this)
}

//Update or Create
func (this *ProjectMaster) Update()  {
	if this.ID != 0 {
		db.Save(this)
	} else {
		db.Create(this)
	}
}

//Retire
func (this ProjectMaster) get() {
	db.Find(ProjectMaster{})

}



