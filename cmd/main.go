package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	db "github.com/Lykeion/lexora-dictionary-service/internal/db"
	grpcserver "github.com/Lykeion/lexora-dictionary-service/internal/grpc"
	pb "github.com/Lykeion/lexora-dictionary-service/internal/grpc/generated"
	"google.golang.org/grpc"
)

func main() {
	dsn, port := getConfigValues()

	// Initialize database connection
	gormDb,err := db.ConnectAndMigrate(dsn); if err != nil{
		fmt.Print(err.Error())
	}
	

	// Initialize services
	referentSvc := &db.ReferentService{DB: *gormDb}
	symbolSvc := &db.SymbolService{DB: *gormDb}
	wordSvc := &db.WordService{DB: *gormDb}
	relationshipSvc := &db.RelationshipService{DB: *gormDb}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	languageServer := &grpcserver.LanguageServer{
		ReferentSvc: referentSvc,
		SymbolSvc:   symbolSvc,
		WordSvc:     wordSvc,
		RelSvc:    relationshipSvc,
	}
	pb.RegisterLanguageServiceServer(grpcServer, languageServer)

	// Start gRPC server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("Starting gRPC server on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped")
}


func getConfigValues() (dsnString string, listenerPort string ){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var port string = ":50051"

	return dsn, port
}
