package grpc

import (
	"log"
	"net"

	"github.com/Tahseen-Zaman/go-grpc/internal/adapters/framework/left/grpc/pb"
	"github.com/Tahseen-Zaman/go-grpc/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	api ports.APIPort
}

func NewAdapter(api ports.APIPort) *Adapter {
	return &Adapter{api: api}
}

func (grpca Adapter) Run() {
	var err error
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	arithmeticServiceServer := grpca
	grpcServer := grpc.NewServer()
	pb.RegisterArithmeticServiceServer(grpcServer, arithmeticServiceServer)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC Server: %v", err)
	}

}
