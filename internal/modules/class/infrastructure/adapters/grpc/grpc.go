package classGRPC

import (
	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	classPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/repositories/postgres"
	memberPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/member/infrastructure/repositories/postgres"
	userGRPCRepo "github.com/zchelalo/sa_class_management/internal/modules/user/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/constants"
	"github.com/zchelalo/sa_class_management/pkg/proto"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
)

type ClassRouter struct {
	useCase *classApplication.ClassUseCases
	proto.UnimplementedClassServiceServer
}

func New(store *db.SQLStore) *ClassRouter {
	classRepository := classPostgresRepo.New(store)

	userClientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userGRPCClient := proto.NewUserServiceClient(userClientConn)
	userRepository := userGRPCRepo.New(userGRPCClient)

	memberRepository := memberPostgresRepo.New(store)

	classUseCases := classApplication.New(classRepository, userRepository, memberRepository)

	return &ClassRouter{
		useCase: classUseCases,
	}
}
