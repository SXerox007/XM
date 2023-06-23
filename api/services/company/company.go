package company

import (
	"github.com/SXerox007/XM/protos/company"
	"google.golang.org/grpc"
)

type Svc struct {
	company.UnimplementedCompanyServiceServer
	adminOrg string
	//TODO: add db
}

// RegisterService with grpc server.
func (h *Svc) RegisterCompanyService(srv *grpc.Server) error {
	company.RegisterCompanyServiceServer(srv, h)
	return nil
}

// TODO: pass db info here
func New(adminOrg string) *Svc {
	return &Svc{adminOrg: adminOrg}
}
