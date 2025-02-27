package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	classGRPC "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/adapters/grpc"
	classPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/repositories/postgres"
	memberApplication "github.com/zchelalo/sa_class_management/internal/modules/member/application"
	memberGRPC "github.com/zchelalo/sa_class_management/internal/modules/member/infrastructure/adapters/grpc"
	memberPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/member/infrastructure/repositories/postgres"
	userGRPCRepo "github.com/zchelalo/sa_class_management/internal/modules/user/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_class_management/internal/server"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/constants"
	"github.com/zchelalo/sa_class_management/pkg/proto"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
	"google.golang.org/grpc"
)

func main() {
	logger := bootstrap.GetLogger()

	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPass,
		config.DBName,
	)
	conn, err := bootstrap.GetInstance("postgres", source)
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	dbStore := db.New(conn)

	classRepository := classPostgresRepo.New(dbStore)
	memberRepository := memberPostgresRepo.New(dbStore)

	services := map[constants.GRPCConstants]string{
		constants.UserMicroserviceDomain: config.UserMicroserviceDomain,
	}

	for name, addr := range services {
		err := bootstrap.InitGRPCClient(addr, name)
		if err != nil {
			logger.Fatalf("cannot init grpc client for %s: %v", name, err)
		}
	}

	userClientConn, err := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	if err != nil {
		logger.Fatalf("cannot get grpc client for %s: %v", constants.UserMicroserviceDomain, err)
	}
	userGRPCClient := proto.NewUserServiceClient(userClientConn)
	userRepository := userGRPCRepo.New(userGRPCClient)

	classUseCases := classApplication.New(classRepository, userRepository, memberRepository)
	memberUseCases := memberApplication.New(classRepository, userRepository, memberRepository)

	classRouter := classGRPC.New(classUseCases)
	memberRouter := memberGRPC.New(memberUseCases)

	server := server.New(config.Port,
		func(s *grpc.Server) { proto.RegisterClassServiceServer(s, classRouter) },
		func(s *grpc.Server) { proto.RegisterMemberServiceServer(s, memberRouter) },
	)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		logger.Println("shutting down gracefully...")
		bootstrap.CloseGRPCClients()
		bootstrap.Close()
		os.Exit(0)
	}()

	server.Start()
}
