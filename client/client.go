package client

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "blowpipe.xyz/blowpipe/blowpiperpc"
	"blowpipe.xyz/blowpipe/constants"
	"google.golang.org/grpc"
)

// const (
// 	address     = "localhost:41415"
// 	defaultName = "world"
// )

// type BPClient interface {
// 	Usage()
// }

type BPClient struct {
	grpc_client pb.BlowpipeClient
	grpc_conn   *grpc.ClientConn
}

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:41415", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

func (c BPClient) Usage() {
	fmt.Printf("blowpipe is a tool for data processing pipelines.\n")
	fmt.Printf("\n")
	fmt.Printf("Usage\n")
	fmt.Printf("\n")
	fmt.Printf("        blowpipe <command> [arguments]\n")
	fmt.Printf("\n")
	fmt.Printf("The server commands are:\n")
	fmt.Printf("        createdb            creates the database.\n")
	fmt.Printf("	    checkdb             verifies the database.\n")
	fmt.Printf("        dropdb              drops the database          [!!CAREFUL!!]\n")
	fmt.Printf("        server              run a blowpipe server\n")
	fmt.Printf("\n")
	fmt.Printf("The client commands are:\n")
	fmt.Printf("\n")
	fmt.Printf("        init              initialises blowpipe configuration and database\n")
	fmt.Printf("        add               add a workflow\n")
	fmt.Printf("        config            (ls|get|set|rm) manages server configuration values\n")
	fmt.Printf("        cancel            cancels a job\n")
	fmt.Printf("        *context          manages client context\n")
	fmt.Printf("        cat (-beta)       print the workflow definition to STDOUT\n")
	fmt.Printf("        disable           disable a workflow\n")
	fmt.Printf("        describe          alter the description of a workflow\n")
	fmt.Printf("        enable            enable a workflow\n")
	fmt.Printf("        history           show revision history for workflow\n")
	fmt.Printf("        ls                list workflows\n")
	fmt.Printf("        log               prints logs\n")
	fmt.Printf("        ps                list jobs\n")
	fmt.Printf("        pause             force a job to pause processing\n")
	fmt.Printf("        rm                delete a workflow\n")
	fmt.Printf("        resume            mark a job to resume processing\n")
	fmt.Printf("        rename            rename a workflow\n")
	fmt.Printf("        *repository       (ls|add|rm) manages repository\n")
	fmt.Printf("        status            print status of the server\n")
	fmt.Printf("        trigger           forces a workflow to commence (starts a job)\n")
	fmt.Printf("        update            update an existing workflow\n")
	fmt.Printf("        version           prints out the version of the client\n")
	fmt.Printf("\n")
	fmt.Printf("Use \"blowpipe help <command>\" for more information about the command.\n")
	fmt.Printf("        (* not implemented)\n\n")

}

func (client BPClient) Createdb() {
	fmt.Println("createdb")
}

func (client BPClient) Checkdb() {
	fmt.Println("checkdb")
}

func (client BPClient) Dropdb() {
	fmt.Println("dropdb")
}

func (client BPClient) Server() {
	fmt.Println("server")
}

func (client BPClient) Init() {
	fmt.Println("init")
}

func (client BPClient) Add() {
	fmt.Println("Add")
}

func (client BPClient) Config() {
	fmt.Println("config")
}

func (client BPClient) Cancel() {
	fmt.Println("Cancel")
}

func (client BPClient) Context() {
	fmt.Println("context")
}

func (client BPClient) Cat() {
	fmt.Println("cat")
}

func (client BPClient) Disable() {
	fmt.Println("disable")
}

func (client BPClient) Describe() {
	fmt.Println("describe")
}

func (client BPClient) Enable() {
	fmt.Println("enable")
}

func (client BPClient) History() {
	fmt.Println("history")
}

func (client BPClient) Ls() {
	fmt.Println("ls")
}

func (client BPClient) Log() {
	fmt.Println("log")
}

func (client BPClient) Pause() {
	fmt.Println("pause")
}

func (client BPClient) Rm() {
	fmt.Println("rm")
}

func (client BPClient) Repository() {
	fmt.Println("repository")
}

func (client BPClient) Resume() {
	fmt.Println("resume")
}

func (client BPClient) Status() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.grpc_client.Status(ctx, &pb.Timestamp{Value: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}

func (client BPClient) Trigger() {
	fmt.Println("trigger")
}

func (client BPClient) Update() {
	fmt.Println("update")
}

func (client BPClient) Version() {
	fmt.Printf("blowpipe version %s\n", constants.VERSION)
}

func (client BPClient) Connect(d time.Duration) {
	var opts []grpc.DialOption

	// opts = append(opts, gprc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(d))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("Failed to dial : %v", err)
	}
	// defer conn.Close()
	client.grpc_client = pb.NewBlowpipeClient(conn)
	client.grpc_conn = conn

}

func (client BPClient) Disconnect() {
	client.grpc_conn.Close()
}

func NewBPClient() *BPClient {
	c := BPClient{}
	return &c
}
