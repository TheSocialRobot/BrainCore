package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/TheSocialRobot/BrainCore/thesocialrobot"
	pb "github.com/TheSocialRobot/BrainCore/thesocialrobot"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func runEventStream(client pb.TheSocialRobotClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.EventStream(ctx)
	if err != nil {
		log.Fatalf("client.EventStream failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.EventStream failed: %v", err)
			}
			log.Printf("Got message %d", in.Id)
			for _, action := range in.Actions {
				switch op := action.Action.(type) {
				case *thesocialrobot.Action_Say:
					log.Printf("delay %d, say %s", action.Delay, op.Say.Text)
				}
			}
		}
	}()
	event := &pb.BodyEvent{Id: 1}
	if err := stream.Send(event); err != nil {
		log.Fatalf("client.EventStream: stream.Send(%v) failed: %v", event, err)
	}
	stream.CloseSend()
	<-waitc
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = "ca_cert.pem"
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
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
	client := pb.NewTheSocialRobotClient(conn)

	runEventStream(client)
}
