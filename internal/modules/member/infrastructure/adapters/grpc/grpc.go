package memberGRPC

// type MemberRouter struct {
// 	useCase *memberApplication.MemberUseCases
// 	memberProto.UnimplementedMemberServiceServer
// }

// func New(store *db.SQLStore) *MemberRouter {
// 	memberRepository := memberPostgresRepo.New(store)

// 	userClientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
// 	userGRPCClient := userProto.NewUserServiceClient(userClientConn)
// 	userRepository := userGRPCRepo.New(userGRPCClient)

// 	memberRepository := memberPostgresRepo.New(store)

// 	memberUseCases := memberApplication.New(memberRepository, userRepository, memberRepository)

// 	return &MemberRouter{
// 		useCase: memberUseCases,
// 	}
// }
