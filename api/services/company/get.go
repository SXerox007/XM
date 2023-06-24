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

func (s *Svc) GetCompany(ctx context.Context, req *company.GetCompanyRequest) (*company.Company, error) {
	// validate the request
	if err := validateGetCompanyRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	result, err := db.GetCompanyDetails(req.GetId())
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

func validateGetCompanyRequest(req *company.GetCompanyRequest) error {
	if err := validation.ValidateStruct(req,
		validation.Field(&req.Id, validation.Required, is.UUID),
	); err != nil {
		log.Println("Error:", err)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
