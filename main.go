package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/getblowpipe/blowpipe/blowpiperpc"
	"github.com/getblowpipe/blowpipe/constants"
	"github.com/simonski/goutils"
	"google.golang.org/grpc"
)

// App is the entrypoint to the whole ballgame
type App struct {
	Client        pb.BlowpipeClient
	Conn          *grpc.ClientConn
	ServerAddress string

	// var (
	// 	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// 	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	// 	serverAddr         = flag.String("server_addr", "localhost:41415", "The server address in the format of host:port")
	// 	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	// )
}

// Main the entrypoiunt to the application - processes input
func (a *App) Main() {

	cli := goutils.NewCLI(os.Args)

	// flag.Parse()
	if len(os.Args) < 2 {
		a.Usage()
	} else {
		// command := os.Args[1]
		command := os.Args[1]
		subCommand := cli.GetStringOrDefault(command, "")
		if command == "server" {
			if subCommand == "createdb" {
				a.Createdb()
			} else if subCommand == "checkdb" {
				a.Checkdb()
			} else if subCommand == "dropdb" {
				a.Dropdb()
			} else if subCommand == "server" {
				a.Server()
			} else {
				// fail

			}
		} else if command == "workflow" {
			if subCommand == "add" {
				a.Add()
			} else if subCommand == "cat" {
				a.Cat()
			} else if subCommand == "trigger" {
				a.Trigger()
			} else if subCommand == "update" {
				a.Update()
			} else if subCommand == "enable" {
				a.Enable()
			} else if subCommand == "disable" {
				a.Disable()
			}

		} else if command == "init" {
			a.Init()

		} else if command == "job" {
			if subCommand == "cancel" {
				a.Cancel()
			} else if subCommand == "pause" {
				a.Pause()
			} else if subCommand == "rm" {
				a.Rm()
			} else if subCommand == "resume" {
				a.Resume()
			}

		} else if command == "config" {
			a.Config()
		} else if command == "context" {
			a.Context()
		} else if command == "describe" {
			a.Describe()
		} else if command == "history" {
			a.History()
		} else if command == "ls" {
			a.Ls()
		} else if command == "log" {
			a.Log()
		} else if command == "repository" {
			a.Repository()
		} else if command == "status" {
			timeout, _ := time.ParseDuration("5s")
			a.Connect(timeout)
			defer a.Disconnect()
			a.Status()
		} else if command == "version" {
			a.Version()
		}
	}

}

func (a *App) Usage() {

	// var (
	// 	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// 	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	// 	serverAddr         = flag.String("server_addr", "localhost:41415", "The server address in the format of host:port")
	// 	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	// )

	// myclient := c.NewBPClient()
	// myclient.Usage()
	// return

	fmt.Printf("blowpipe is a tool for data processing pipelines.\n")
	fmt.Printf("\n")
	fmt.Printf("Usage\n")
	fmt.Printf("\n")
	fmt.Printf("        blowpipe <command> [arguments]\n")
	fmt.Printf("\n")
	fmt.Printf("The server commands are:\n\n")
	fmt.Printf("        createdb          creates the database.\n")
	fmt.Printf("        checkdb           verifies the database.\n")
	fmt.Printf("        dropdb            drops the database          [!!CAREFUL!!]\n")
	fmt.Printf("        server            run a blowpipe server\n")
	fmt.Printf("\n")
	fmt.Printf("The client commands are:\n")
	fmt.Printf("\n")
	fmt.Printf("        init              initialises blowpipe configuration and database\n")
	fmt.Printf("\n")
	fmt.Printf("The workflow commands are:\n")
	fmt.Printf("\n")
	fmt.Printf("        add               add a workflow\n")
	fmt.Printf("        ls                list workflows\n")
	fmt.Printf("        cat (-beta)       print the workflow definition to STDOUT\n")
	fmt.Printf("        disable           disable a workflow\n")
	fmt.Printf("        rm                delete a workflow\n")
	fmt.Printf("        rename            rename a workflow\n")
	fmt.Printf("        describe          alter the description of a workflow\n")
	fmt.Printf("        enable            enable a workflow\n")
	fmt.Printf("        history           show revision history for workflow\n")
	fmt.Printf("        trigger           forces a workflow to commence (starts a job)\n")
	fmt.Printf("        update            update an existing workflow\n")
	fmt.Printf("\n")
	fmt.Printf("The job commands are:\n")
	fmt.Printf("\n")
	fmt.Printf("        ps                list jobs\n")
	fmt.Printf("        pause             force a job to pause processing\n")
	fmt.Printf("        resume            mark a job to resume processing\n")
	fmt.Printf("        cancel            cancels a job\n")
	fmt.Printf("\n")
	fmt.Printf("The other commands are:\n")
	fmt.Printf("\n")

	fmt.Printf("        config            (ls|get|set|rm) manages server configuration values\n")
	fmt.Printf("        *context          manages client context\n")
	fmt.Printf("        log               prints logs\n")
	fmt.Printf("        *repository       (ls|add|rm) manages repository\n")
	fmt.Printf("        status            print status of the server\n")
	fmt.Printf("        version           prints out the version of the client\n")
	fmt.Printf("\n")
	fmt.Printf("Use \"blowpipe help <command>\" for more information about the command.\n")
	fmt.Printf("        (* not implemented)\n\n")

}

// Createdb creates the database
func (a *App) Createdb() {
	fmt.Println("createdb")
}

// Checkdb verifies the DB exists and checks its status
func (a *App) Checkdb() {
	fmt.Println("checkdb")
}

// Dropdb deletes the database - WARNING this is destructive and permanent
func (a *App) Dropdb() {
	fmt.Println("dropdb")
}

// Server runs the blowpipe server process
func (a *App) Server() {
	fmt.Println("server")
}

// Init initialises the blowpipe.cfg and .db files
// The autodiscover logic will attempt to find in order
// A CLI parameter, the BLOWPIPE_HOME directory or the current directory
func (a *App) Init() {
	fmt.Println("init")
}

func (a *App) Add() {
	fmt.Println("Add")
}

func (a *App) Config() {
	fmt.Println("config")
}

func (a *App) Cancel() {
	fmt.Println("Cancel")
}

func (a *App) Context() {
	fmt.Println("context")
}

func (a *App) Cat() {
	fmt.Println("cat")
}

func (a *App) Disable() {
	fmt.Println("disable")
}

func (a *App) Describe() {
	fmt.Println("describe")
}

func (a *App) Enable() {
	fmt.Println("enable")
}

func (a *App) History() {
	fmt.Println("history")
}

func (a *App) Ls() {
	fmt.Println("ls")
}

func (a *App) Log() {
	fmt.Println("log")
}

func (a *App) Pause() {
	fmt.Println("pause")
}

func (a *App) Rm() {
	fmt.Println("rm")
}

func (a *App) Repository() {
	fmt.Println("repository")
}

func (a *App) Resume() {
	fmt.Println("resume")
}

func (a *App) Status() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := a.Client.Status(ctx, &pb.Timestamp{Value: 5})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}

func (a *App) Trigger() {
	fmt.Println("trigger")
}

func (a *App) Update() {
	fmt.Println("update")
}

func (a *App) Version() {
	fmt.Printf("blowpipe version %s\n", constants.VERSION)
}

func (a *App) Connect(d time.Duration) {
	var opts []grpc.DialOption

	// opts = append(opts, gprc.WithInsecure())
	opts = append(opts, grpc.WithTimeout(d))
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithInsecure())

	ServerAddress := "localhost:41415"
	// conn, err := grpc.Dial(a.ServerAddress, opts...)
	conn, err := grpc.Dial(ServerAddress, opts...)
	if err != nil {
		log.Fatalf("Failed to dial : %v, %v", ServerAddress, err)
	}
	a.Client = pb.NewBlowpipeClient(conn)
	a.Conn = conn

	// TODO I don't understand in this contet why defer is triggering
	// there is a reference in the App however it still closes.
	// So right now I leave it to the client to call disconnect
	// defer conn.Close()
	log.Printf("Connected!\n")

}

func (a *App) Disconnect() {
	a.Conn.Close()
}

func main() {
	app := &App{}
	app.Main()
}
