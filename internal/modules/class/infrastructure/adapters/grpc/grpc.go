package classGRPC

import (
	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	"github.com/zchelalo/sa_class_management/pkg/proto"
)

type ClassRouter struct {
	useCase *classApplication.ClassUseCases
	proto.UnimplementedClassServiceServer
}

func New(classUseCase *classApplication.ClassUseCases) *ClassRouter {
	return &ClassRouter{
		useCase: classUseCase,
	}
}
