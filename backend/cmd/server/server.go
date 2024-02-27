package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"

	pb "github.com/Rhaqim/grpc-python-go/internal/infrastructure/backend/recommendationpb"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type recommenderServer struct {
	pb.UnimplementedRecommenderServer
}

func (s *recommenderServer) GetRecommendations(ctx context.Context, in *pb.RecommendationRequest) (*pb.RecommendationResponse, error) {

	recommendedItems := []string{"item1", "item2", "item3"}

	return &pb.RecommendationResponse{RecommendedItems: recommendedItems}, nil
}

func NewServer() *recommenderServer {
	return &recommenderServer{}
}

func RunServer() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = data.Path("x509/server_cert.pem")
		}
		if *keyFile == "" {
			*keyFile = data.Path("x509/server_key.pem")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRecommenderServer(grpcServer, NewServer())

	grpcServer.Serve(lis)
}
