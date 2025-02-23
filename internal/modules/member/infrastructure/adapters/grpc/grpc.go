package memberGRPC

import (
	classPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/class/infrastructure/repositories/postgres"
	memberApplication "github.com/zchelalo/sa_class_management/internal/modules/member/application"
	memberPostgresRepo "github.com/zchelalo/sa_class_management/internal/modules/member/infrastructure/repositories/postgres"
	userGRPCRepo "github.com/zchelalo/sa_class_management/internal/modules/user/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/constants"
	"github.com/zchelalo/sa_class_management/pkg/proto"
	"github.com/zchelalo/sa_class_management/pkg/sqlc/db"
)

type MemberRouter struct {
	useCase *memberApplication.MemberUseCases
	proto.UnimplementedMemberServiceServer
}

func New(store *db.SQLStore) *MemberRouter {
	classRepository := classPostgresRepo.New(store)

	memberRepository := memberPostgresRepo.New(store)

	userClientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userGRPCClient := proto.NewUserServiceClient(userClientConn)
	userRepository := userGRPCRepo.New(userGRPCClient)

	memberUseCases := memberApplication.New(classRepository, userRepository, memberRepository)

	return &MemberRouter{
		useCase: memberUseCases,
	}
}
