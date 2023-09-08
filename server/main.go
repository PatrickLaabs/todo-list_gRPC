package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/PatrickLaabs/todo-list_grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	port = flag.Int("port", 50055, "The server port")
)

type server struct {
	pb.UnimplementedGreeterServer
	db *gorm.DB
}

type Name struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())

	name := &Name{Name: in.GetName()}
	if err := s.db.Create(name).Error; err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Failed to insert name into database: %v", err))
	}
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	dbPWD := os.Getenv("DBPWD")
	dbURL := "postgresql://patrick:" + dbPWD + "@ornate-dolphin-3264.8nj.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db.AutoMigrate(&Name{})

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{db: db})

	log.Printf("Starting server on port %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
