package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/SXerox007/XM/base/environment"
	"github.com/SXerox007/XM/protos/company"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

// JWTRequest represents the request body for JWT token generation.
type JWTRequest struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

// JWTResponse represents the response body containing the generated JWT token.
type JWTResponse struct {
	Token string `json:"token"`
}

var (
	authpoint     = flag.String("auth_end_points", "localhost:50051", "expose xm end point")
	authEndpointD = flag.String("auth_end_point", "xm-api:50051", "XM gRPC endpoint")
)

func ExposePoint(address string, middleware func(http.Handler) http.Handler, opts ...runtime.ServeMuxOption) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ap := authpoint
	if isDocker := os.Getenv("IS_DOCKER"); isDocker == "true" {
		log.Println("Running inside Docker container")
		// Perform any necessary configuration specific to Docker environment
		ap = authEndpointD
	}

	log.Println("authEndPoint", ap)

	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	// Register gRPC gateway
	err := company.RegisterCompanyServiceHandlerFromEndpoint(ctx, mux, *ap, dialOpts)
	if err != nil {
		return err
	}

	// Apply the authentication middleware
	handler := middleware(mux)

	grpcMux := http.NewServeMux()
	grpcMux.Handle("/", handler)

	log.Println("Starting Endpoint Exposed Server:", address)
	return http.ListenAndServe(address, grpcMux)
}

func init() {
	env := environment.GetRestEnv()
	port := environment.GetRestPort()

	address := env + port
	if err := ExposePoint(address, authMiddleware); err != nil {
		log.Fatal("Error in serve", err)
	}
}

func main() {
	flag.Parse()
}
