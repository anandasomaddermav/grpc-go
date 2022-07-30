package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"github.com/golang/protobuf/proto"
	pbo1pm "mavenir.com/o1pm/o1pmstream"
)

var (
	port       = flag.Int("port", 50051, "The server port")
)

type o1pmServer struct {
	pbo1pm.UnimplementedPMStreamServer

}

// StreamPMData streams data from the client side towards the o1pm server
//
// It gets a stream of pdsus, and responds with empty object

func (s *o1pmServer) StreamPMData(stream pbo1pm.PMStream_StreamPMDataServer) error {
	startTime := time.Now()
	for {
		pdsus, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			return stream.SendAndClose(&pbo1pm.emptypb.Empty)
		}
		if err != nil {
			return err
		}

	}
}

func newServer() *o1pmServer {
	s := &o1pmServer{}
	//s.loadData(*jsonDBFile)
	return s
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterPMStreamServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
