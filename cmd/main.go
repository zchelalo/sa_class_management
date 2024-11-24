package main

import (
	"fmt"

	"github.com/zchelalo/sa_class_management/internal/server"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/constants"
)

func main() {
	bootstrap.InitLogger()
	logger := bootstrap.GetLogger()

	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)
	err = bootstrap.InitInstance("postgres", source)
	if err != nil {
		logger.Fatal("cannot connect to db:", err)
	}

	services := map[constants.GRPCConstants]string{
		constants.UserMicroserviceDomain: config.UserMicroserviceDomain,
	}

	for name, addr := range services {
		err := bootstrap.InitGRPCClient(addr, name)
		if err != nil {
			logger.Fatalf("cannot init grpc client for %s: %v", name, err)
		}
	}

	server.Start()
}
