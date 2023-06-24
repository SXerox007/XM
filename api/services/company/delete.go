package company

import (
	"context"

	db "github.com/SXerox007/XM/api/storage/postgres"
	"github.com/SXerox007/XM/protos/company"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Svc) DeleteCompany(ctx context.Context, req *company.DeleteCompanyRequest) (*emptypb.Empty, error) {
	// validate the request
	if err := validateGetCompanyRequest(&company.GetCompanyRequest{
		Id: req.Id,
	}); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err := db.RemoveCompanyDetails(req.Id)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
