package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	db "github.com/SXerox007/XM/base/postgres"

	"github.com/SXerox007/XM/base/environment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cp "github.com/SXerox007/XM/api/services/company"
)

func init() {
	pg := db.InitPG()
	defer pg.Close()
}

func main() {
	env := environment.GetEnv()
	port := environment.GetPort()
	ServerSetup(env, port)
}

// ServerSetup sets up and starts the gRPC server
func ServerSetup(env, port string) {
	listener, err := net.Listen(env, port)
	if err != nil {
		log.Println("Error in server start:", err)
		return
	}

	// Create a new gRPC server
	srv := grpc.NewServer()

	// Register reflection on gRPC server for debugging purposes
	reflection.Register(srv)

	// Register the services
	setupServices(srv)

	// Start the server in a goroutine
	go func() {
		fmt.Println("Server start on", env+port)
		if err := srv.Serve(listener); err != nil {
			log.Println("Error in Serve:", err)
			return
		}
	}()

	// Make a channel that will wait for server to close or interrupt
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block the code until server is interrupted or closed
	<-ch
	log.Println("Exiting the server with:", os.Interrupt)
}

// setupServices registers all the services with the gRPC server
func setupServices(srv *grpc.Server) {
	rcs := cp.New("xm")
	rcs.RegisterCompanyService(srv)
}
