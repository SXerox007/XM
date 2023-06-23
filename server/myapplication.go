package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	db "github.com/SXerox007/XM/base/postgres"

	"github.com/SXerox007/XM/base/environment"
	"github.com/SXerox007/XM/base/server"
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

// brain setup
func ServerSetup(env, port string) {
	listner, err := net.Listen(env, port)
	if err != nil {
		log.Println("Error in server start:", err)
		return
	}
	//Create the New gRPC Server
	srv := server.CreateNewgRPCServer(false, nil)
	//Register reflection on gRPC server
	reflection.Register(srv)
	// all the rpc services
	setupServices(srv)

	go func() {
		fmt.Println("Server start on", env+port)
		if err := srv.Serve(listner); err != nil {
			log.Println("Error in Serve:", err)
			return
		}
	}()
	//make a channnel that will wait for server to close or interrupt
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// this will block the code while server
	<-ch
	log.Println("Exit the Server with:", os.Interrupt)
}

// All the services
func setupServices(srv *grpc.Server) {

	rcs := cp.New("xm")
	rcs.RegisterCompanyService(srv)
}
