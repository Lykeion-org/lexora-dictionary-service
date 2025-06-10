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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"google.golang.org/grpc"
)

func main() {
	// Database configuration
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Initialize database connection
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto migrate database schema
	if err := gormDB.AutoMigrate(
		&db.Referent{},
		&db.Symbol{},
		&db.Word{},
	); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize services
	referentSvc := &db.ReferentService{DB: *gormDB}
	symbolSvc := &db.SymbolService{DB: *gormDB}
	wordSvc := &db.WordService{DB: *gormDB}

	// Create gRPC server
	grpcServer := grpc.NewServer()
	languageServer := &grpcserver.LanguageServer{
		ReferentSvc: referentSvc,
		SymbolSvc:   symbolSvc,
		WordSvc:     wordSvc,
	}
	pb.RegisterLanguageServiceServer(grpcServer, languageServer)

	// Start gRPC server
	listener, err := net.Listen("tcp", ":50051")
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
