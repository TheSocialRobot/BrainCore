package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials"

	pb "github.com/TheSocialRobot/BrainCore/thesocialrobot"
)

var (
	tls      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50051, "The server port")
)

type theSocialRobotServer struct {
	pb.UnimplementedTheSocialRobotServer
}

func (s *theSocialRobotServer) EventStream(stream pb.TheSocialRobot_EventStreamServer) error {
	for {
		// TODO handle events from the client
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received event, sending one command")
		// respond with a single command
		// TODO eventually we'll decouple receiving events from sending commands
		command := &pb.BodyCommand{
			Id:      1,
			Actions: []*pb.Action{{Delay: 0, Action: &pb.Action_Say{Say: &pb.Say{Text: "Hello World"}}}},
		}
		stream.Send(command)
	}
}

func newServer() *theSocialRobotServer {
	return new(theSocialRobotServer)
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			// TODO do this properly / create this file
			*certFile = "server_cert.pem"
		}
		if *keyFile == "" {
			// TODO do this properly / create this file
			*keyFile = "server_key.pem"
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTheSocialRobotServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
