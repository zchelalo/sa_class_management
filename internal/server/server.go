package server

import (
	"context"
	"fmt"
	"net"

	_ "github.com/lib/pq"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
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

	// store := db.New(conn)
	_ = db.New(conn)

	// ctx := context.Background()
	_ = context.Background()

	// userRouter := userInfrastructure.NewUserRouter(userStore, ctx)
	config := bootstrap.GetConfig()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		logger.Fatal("cannot listen:", err)
	}

	server := grpc.NewServer()
	// userProto.RegisterUserServiceServer(server, userRouter)

	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		logger.Fatalf("Error serving: %s", err.Error())
	}

	logger.Println("Server running on port:", config.Port)
}
