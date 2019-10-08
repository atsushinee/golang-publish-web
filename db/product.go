package db

import (
	"fmt"
	"github.com/atsushinee/golang-publish-web/models"
	"github.com/jeanphorn/log4go"
)

func GetProductionsByProjectId(projectId int64) []*models.Product {
	product := &models.Product{
		ProjectId: projectId,
	}
	var products []*models.Product
	rows, err := x.Rows(product)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.Product)
		err = rows.Scan(p)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	return products
}

func CreateProduct(projectId int64, productName string) error {
	_, err := x.InsertOne(&models.Product{
		ProjectId: projectId,
		Name:      productName,
	})
	if err != nil {
		log4go.Error(fmt.Sprintf("CreateProduct error : %s", err))
	}
	return err
}
