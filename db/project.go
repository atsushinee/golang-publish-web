package db

import (
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/jeanphorn/log4go"
)

func CreateProject(name string) error {
	project := &models.Project{
		Name: name,
	}
	_, err := x.InsertOne(project)
	if err != nil {
		log4go.Error("CreateUser error:", err)
	}
	return err
}

func GetAllProjects() []*models.Project {
	var projects []*models.Project
	project := new(models.Project)
	rows, err := x.Rows(project)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.Project)
		err = rows.Scan(p)
		if err != nil {
			panic(err)
		}
		projects = append(projects, p)
	}
	return projects
}

func GetAllProjectWithProducts() []*models.Project {
	var projects []*models.Project
	project := new(models.Project)
	rows, err := x.Rows(project)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.Project)
		err = rows.Scan(p)
		if err != nil {
			panic(err)
		}
		p.Products = GetProductionsByProjectId(p.Id)
		projects = append(projects, p)
	}
	return projects
}

func GetAllProjectWithDocs() []*models.Project {
	var projects []*models.Project
	project := new(models.Project)
	rows, err := x.Rows(project)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.Project)
		err = rows.Scan(p)
		if err != nil {
			panic(err)
		}
		docs, err := GetAllDocsByPid(p.Id)
		if err != nil {
			panic(err)
		}
		p.Docs = docs
		projects = append(projects, p)
	}
	return projects
}
