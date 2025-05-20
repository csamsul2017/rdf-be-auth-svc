package main

import (
	"log"
	"net"
	"os"

	// "github.com/csamsul2017/rdf-be-auth-svc.git/delivery/grpc"
	// "github.com/csamsul2017/rdf-be-auth-svc.git/exceptions"
	"github.com/csamsul2017/rdf-be-auth-svc.git/helper"
	"github.com/csamsul2017/rdf-be-auth-svc.git/infrastructure"

	// "github.com/csamsul2017/rdf-be-auth-svc.git/proto"
	// "github.com/csamsul2017/rdf-be-auth-svc.git/repository"
	// "github.com/csamsul2017/rdf-be-auth-svc.git/usecase"
	"github.com/joho/godotenv"

	gogrpc "google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	db := infrastructure.ConnectDB()

	// Run migrations
	if err := helper.RunMigrations(db); err != nil {
		log.Fatal("Migrations failed: ", err)
	}

	// Run seeds
	if err := helper.RunSeeds(db); err != nil {
		log.Fatal("Seeding failed: ", err)
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	gRPCErrorTranslator := exceptions.NewGRPCErrorTranslator()
	validator := helper.NewValidator()
	grpcValidator := helper.NewGRPCValidator(validator, gRPCErrorTranslator)
	scopeRepo := repository.NewScopeRepository(db)
	roleScopeRepo := repository.NewRoleScopeRepository(db)
	scopeUC := usecase.NewScopeUsecase(scopeRepo, roleScopeRepo)
	scopeGRPC := grpc.NewScopeGRPC(scopeUC, grpcValidator, gRPCErrorTranslator)

	gRPCServer := gogrpc.NewServer()
	proto.RegisterScopeServiceServer(gRPCServer, scopeGRPC)

	log.Println("gRPC server is running on port " + port)
	err = gRPCServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve gRPC server over port "+port+": %v", err)
	}
}