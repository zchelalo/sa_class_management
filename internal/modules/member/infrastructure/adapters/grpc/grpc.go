package memberGRPC

import (
	memberApplication "github.com/zchelalo/sa_class_management/internal/modules/member/application"
	"github.com/zchelalo/sa_class_management/pkg/proto"
)

type MemberRouter struct {
	useCase *memberApplication.MemberUseCases
	proto.UnimplementedMemberServiceServer
}

func New(memberUseCase *memberApplication.MemberUseCases) *MemberRouter {
	return &MemberRouter{
		useCase: memberUseCase,
	}
}
