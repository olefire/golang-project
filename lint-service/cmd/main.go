package main

import (
	"google.golang.org/grpc"
	gapi "lint-service/internal/gapi/linters"
	"lint-service/internal/linters"
	"lint-service/internal/linters/python/metrics"
	"lint-service/internal/linters/python/pylint"
	"lint-service/internal/services/linter"
	"lint-service/pkg/protos/gen"
	"log"
	"net"
)

var (
	linterService *linter.Service
)

func main() {
	pyLint := pylint.Linter{}
	pyMetrics := metrics.Linter{}

	l := []linters.Linter{&pyLint, &pyMetrics}
	linterService = linter.NewClient(l)

	startGrpcServer()
}

func startGrpcServer() {
	linterServer := gapi.NewGrpcServer(gapi.LintingService{LinterManagement: linterService})
	grpcServer := grpc.NewServer()

	gen.RegisterLintingServiceServer(grpcServer, linterServer)

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
