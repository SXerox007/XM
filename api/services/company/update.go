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

func (s *Svc) UpdateCompany(ctx context.Context, req *company.UpdateCompanyRequest) (*company.Company, error) {
	// validate the request
	if err := validateUpdateCompanyRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := db.UpdateCompanyDetails(db.Company{
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Employees:   req.GetEmployees(),
		Registered:  req.GetRegistered(),
		Type:        req.GetType().String(),
		Id:          req.GetId(),
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
func validateUpdateCompanyRequest(req *company.UpdateCompanyRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, is.UUID),
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
