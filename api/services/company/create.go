package company

import (
	"context"
	"log"

	db "github.com/SXerox007/XM/api/storage/postgres"
	"github.com/SXerox007/XM/protos/company"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Svc) CreateCompany(ctx context.Context, req *company.CreateCompanyRequest) (*company.Company, error) {
	log.Println("Here at CreateCompany")
	// validate the request
	if err := validate(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := db.CreateCompanyRow(db.Company{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Employees:   req.GetEmployees(),
		Registered:  req.GetRegistered(),
		Type:        req.GetType().String(),
	})

	if err != nil {
		return nil, err
	}

	return &company.Company{
		Id:          result.Id,
		Name:        result.Name,
		Description: result.Description,
		Employees:   result.Employees,
		Registered:  result.Registered,
		//Type:        result.Type,
	}, nil

}

// validate the request
func validate(req *company.CreateCompanyRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 15)),
		validation.Field(&req.Description, validation.Length(0, 3000)),
		validation.Field(&req.Employees, validation.Required, is.Digit),
		validation.Field(&req.Registered, validation.Required),
		validation.Field(&req.Type, validation.Required, validation.In("Corporations", "NonProfit", "Cooperative", "Sole Proprietorship")),
	); err != nil {
		log.Println("Error in validate calculateOrderValue:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
