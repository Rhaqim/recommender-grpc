package client

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/Rhaqim/grpc-python-go/internal/infrastructure/backend/recommendationpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

// printFeature gets the feature for the given point.
func getrecommendations(client pb.RecommenderClient, request *pb.RecommendationRequest) {
	log.Printf("Getting recommendations for user_id: %s with %d", request.UserId, request.NumRecommendations)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	recommendations, err := client.GetRecommendations(ctx, request)
	if err != nil {
		log.Fatalf("%v.GetRecommendations(_) = _, %v: ", client, err)
	}
	log.Println(recommendations)
}

func MakeRequest() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewRecommenderClient(conn)

	// Looking for a valid feature
	getrecommendations(client, &pb.RecommendationRequest{UserId: "1", NumRecommendations: 5})

}
