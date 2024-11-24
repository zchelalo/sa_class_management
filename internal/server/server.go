package server

import (
	"fmt"
	"net"

	_ "github.com/lib/pq"
	classGRPC "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/adapters/grpc"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	classProto "github.com/zchelalo/sa_class_management/pkg/proto/class"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {
	logger := bootstrap.GetLogger()

	conn, err := bootstrap.GetInstance()
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	store := db.New(conn)

	config := bootstrap.GetConfig()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Fatal("cannot listen:", err)
	}

	server := grpc.NewServer()

	classRouter := classGRPC.New(store)
	classProto.RegisterClassServiceServer(server, classRouter)

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		logger.Fatalf("Error serving: %s", err.Error())
	}

	logger.Println("Server running on port:", config.Port)
}
