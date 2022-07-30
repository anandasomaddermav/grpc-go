package o1pmclient

import (
		"context"
		"flag"
		"io"
		"log"
		"time"
		"google.golang.org/grpc"
		pbo1pm "mavenir.com/o1pm/o1pmstream"
       )

var (
		serverAddr = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
		jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
    )

// loadData loads data from a exampleData or file (P-2).
func (client pbo1pm.PMStreamClient) loadData() {
	var data []byte
		data = exampleData
		if err := json.Unmarshal(data, &client.PMStreamClient); err != nil {
			log.Fatalf("Failed to load data: %v", err)
		}
}

// sendStreamRecord sends a stream record to server and expects to get an Empty response from server.
func sendStreamPMData(client pbo1pm.PMStreamClient) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		stream, err := client.StreamPMData(ctx)
		if err != nil {
			log.Fatalf("client.StreamPMData failed: %v", err)
		}
	for _, point := range points {
		if err := stream.Send(point); err != nil {
			log.Fatalf("client.StreamPMData: stream.Send failed: %v", err)
		}
	}
	reply, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalf("client.StreamPMData failed: %v", err)
		}
	log.Printf("Route summary: %v", reply)
}

func main() {
	var opts []grpc.DialOption

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewPMStreamClient(conn)

	// StreamPMData
	sendStreamPMData(client)

}

// exampleData is a copy of testdata/route_guide_db.json. It's to avoid
// specifying file path with `go run`.
var exampleData = []byte(`[{
	"location": {
	"latitude": 407838351,
	"longitude": -746143763
	},
	"name": "Patriots Path, Mendham, NJ 07945, USA"
	}, {
	"location": {
	"latitude": 408122808,
	"longitude": -743999179
	},
	"name": "101 New Jersey 10, Whippany, NJ 07981, USA"
	}
	]`)

