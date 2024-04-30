package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-zookeeper/zk"
	pb "github.com/portanic/portanic/api/proto"
	"google.golang.org/grpc"
)

func connectZK() *zk.Conn {
	conn, _, err := zk.Connect([]string{"zookeeper:2181"}, time.Second*10)
	if err != nil {
		panic(err)
	}
	return conn
}

func registerService(conn *zk.Conn, path string, data []byte) error {
	flags := int32(zk.FlagEphemeral)
	acl := zk.WorldACL(zk.PermAll)
	_, err := conn.Create(path, data, flags, acl)
	if err != nil {
		return err
	}
	fmt.Println("Service registered:", path)
	return nil
}

type server struct {
	pb.UnimplementedPluginServiceServer
}

func (s *server) GetBlock(ctx context.Context, req *pb.BlockRequest) (*pb.BlockResponse, error) {
	return &pb.BlockResponse{
		Success:     true,
		HtmlContent: "<div>Hello, Block!</div>",
	}, nil
}

func main() {
	zkConn := connectZK()
	defer zkConn.Close()
	time.Sleep(2 * time.Second)
	createPersistentNodeIfNeeded(zkConn, "/services")
	err := registerService(zkConn, "/services/example", []byte("1.0.0.0:50051"))
	if err != nil {
		fmt.Println("Failed to register service:", err)
	}
	fmt.Println("Done registering")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPluginServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func createPersistentNodeIfNeeded(conn *zk.Conn, path string) {
	exists, _, err := conn.Exists(path)
	if err != nil {
		fmt.Printf("Error checking existence of node %s: %v\n", path, err)
		return
	}
	if !exists {
		// Node does not exist, create it
		flags := int32(0) // Persistent node
		acl := zk.WorldACL(zk.PermAll)
		_, err := conn.Create(path, []byte{}, flags, acl)
		if err != nil {
			fmt.Printf("Failed to create node %s: %v\n", path, err)
			return
		}
		fmt.Printf("Node created: %s\n", path)
	} else {
		fmt.Printf("Node already exists: %s\n", path)
	}
}
