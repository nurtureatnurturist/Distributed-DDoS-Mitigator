package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 5999
// Hash 8088
// Hash 2094
// Hash 3804
// Hash 8278
// Hash 9058
// Hash 5045
// Hash 5671
// Hash 8216
// Hash 5630
// Hash 3676
// Hash 5168
// Hash 9333
// Hash 4128
// Hash 2250
// Hash 3667
// Hash 2823
// Hash 5470
// Hash 7157
// Hash 5580
// Hash 9549
// Hash 9480
// Hash 8950
// Hash 4494
// Hash 4227
// Hash 2830
// Hash 5948
// Hash 8128
// Hash 7591
// Hash 1750
// Hash 9186