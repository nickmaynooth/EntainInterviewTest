package main

import (
	"database/sql"
	"flag"
	"log"
	"net"

	"git.neds.sh/matty/entain/sports/db"
	"git.neds.sh/matty/entain/sports/proto/sports"
	"git.neds.sh/matty/entain/sports/service"
	"google.golang.org/grpc"
)

var (
	grpcEndpoint = flag.String("grpc-endpoint", "localhost:6000", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running grpc server: %s\n", err)
	}
}

func run() error {
	conn, err := net.Listen("tcp", ":6000")
	if err != nil {
		return err
	}

	eventsDB, err := sql.Open("sqlite3", "./db/sports.db")
	if err != nil {
		return err
	}

	eventsRepo := db.NewEventsRepo(eventsDB)
	if err := eventsRepo.Init(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	sports.RegisterSportsServer(
		grpcServer,
		service.NewSportsService(
			eventsRepo,
		),
	)

	log.Printf("gRPC server listening on: %s\n", *grpcEndpoint)

	if err := grpcServer.Serve(conn); err != nil {
		return err
	}

	return nil
}
