package postgres

import db "github.com/SXerox007/XM/base/postgres"

type Company struct {
	TableName   struct{} `sql:"companies" json:"-"`
	Id          string   `param:"id" json:"id" sql:"id"`
	Name        string   `param:"name" json:"name" sql:"name"`
	Description string   `param:"description" json:"description" sql:"description"`
	Employees   int32    `param:"employees" json:"employees" sql:"employees"`
	Registered  bool     `param:"registered" json:"registered" sql:"registered"`
	Type        string   `param:"type" json:"type" sql:"type"`
}

// store company details
func CreateCompanyRow(data Company) (Company, error) {
	err := db.PG.Create(&data)
	return data, err
}

// update company details
func UpdateCompanyDetails(data Company) (Company, error) {
	return data, nil
}
