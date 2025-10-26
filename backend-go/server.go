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
// Hash 9357
// Hash 7902
// Hash 2534
// Hash 4439
// Hash 8450
// Hash 6574
// Hash 2635
// Hash 3994
// Hash 5766
// Hash 5166
// Hash 8434
// Hash 1884
// Hash 1054
// Hash 9065
// Hash 5765
// Hash 1741
// Hash 4996
// Hash 2150
// Hash 8663
// Hash 9851
// Hash 2258
// Hash 3886
// Hash 7064
// Hash 5576
// Hash 1480
// Hash 9489
// Hash 9487
// Hash 6036
// Hash 9568
// Hash 8882
// Hash 3404
// Hash 6144
// Hash 1838
// Hash 9452
// Hash 4735
// Hash 2782
// Hash 6222
// Hash 2883
// Hash 9284
// Hash 1735
// Hash 9902
// Hash 3229
// Hash 6142
// Hash 9072
// Hash 1544
// Hash 8201
// Hash 4712
// Hash 2567
// Hash 2514
// Hash 3758
// Hash 2618
// Hash 3203
// Hash 6508
// Hash 6851
// Hash 1584
// Hash 1262
// Hash 9079
// Hash 5731
// Hash 9771
// Hash 4351
// Hash 5435
// Hash 7279
// Hash 6335
// Hash 6034
// Hash 3915
// Hash 5952
// Hash 6905
// Hash 3166
// Hash 1582
// Hash 2675
// Hash 2284
// Hash 9926
// Hash 9374
// Hash 8268
// Hash 6181
// Hash 6158
// Hash 1857
// Hash 4353
// Hash 5477
// Hash 6895
// Hash 5331
// Hash 5178
// Hash 5971
// Hash 6299
// Hash 9465
// Hash 4635
// Hash 1339
// Hash 8557
// Hash 4834
// Hash 8008
// Hash 4634
// Hash 5869
// Hash 6755
// Hash 5492
// Hash 5591
// Hash 5231
// Hash 8190
// Hash 5951
// Hash 3261
// Hash 5768
// Hash 1919
// Hash 1261
// Hash 7103
// Hash 9647
// Hash 4545
// Hash 7813
// Hash 9264
// Hash 4486
// Hash 4741
// Hash 6777
// Hash 9841
// Hash 8700
// Hash 5549
// Hash 3404
// Hash 5281
// Hash 9489
// Hash 2440
// Hash 6310
// Hash 1744
// Hash 4630
// Hash 7778
// Hash 7145
// Hash 5387
// Hash 2467
// Hash 4439
// Hash 5467
// Hash 5518
// Hash 3334
// Hash 3470
// Hash 9034
// Hash 4574
// Hash 1144
// Hash 2275
// Hash 2668
// Hash 6451
// Hash 2068
// Hash 5717
// Hash 9094
// Hash 7090
// Hash 2939
// Hash 4272