package postgres

import (
	"errors"

	db "github.com/SXerox007/XM/base/postgres"
)

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

// get details
func GetCompanyDetails(id string) (data Company, err error) {
	err = db.PG.Model(&Company{}).Where("id = ?", id).Order("uts desc").Limit(1).Select(&data)
	return
}

// update company details
func UpdateCompanyDetails(data Company) (Company, error) {
	if compData, err := GetCompanyDetails(data.Id); compData.Id != "" && err == nil {
		// do update
		_, err := db.PG.Model(&Company{}).Set("name = ?", data.Name).Set("description = ?", data.Description).Set("employees = ?", data.Employees).Set("registered = ?", data.Registered).Set("type = ?", data.Type).
			Where("id = ?", data.Id).Update()
		return data, err
	}
	return data, errors.New("invalid data")
}

// remove company details
func RemoveCompanyDetails(id string) error {
	_, err := db.PG.Model(&Company{}).Where("id = ?", id).Delete()
	return err
}
