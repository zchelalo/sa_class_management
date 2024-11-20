package classApplication

import (
	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
)

type ClassUseCases struct {
	classRepository classDomain.ClassRepository
	UserRepository  userDomain.UserRepository
}

func New(classRepository classDomain.ClassRepository, userRepository userDomain.UserRepository) *ClassUseCases {
	return &ClassUseCases{
		classRepository: classRepository,
		UserRepository:  userRepository,
	}
}
