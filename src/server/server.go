package server

import (
	"context"
	"log"
	"net"

	protos "github.com/datastx/IceBridge/src/protos/pbs"
	"google.golang.org/grpc"
)

type healthServer struct {
	protos.UnimplementedHealthServer
}

func (s *healthServer) Check(ctx context.Context, req *protos.HealthCheckRequest) (*protos.HealthCheckResponse, error) {
	status := protos.HealthCheckResponse_SERVING
	// Logic to determine server health status
	return &protos.HealthCheckResponse{Status: status}, nil
}

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterHealthServer(s, &healthServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
