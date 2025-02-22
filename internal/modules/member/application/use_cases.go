package memberApplication

import (
	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
)

type MemberUseCases struct {
	classRepository  classDomain.ClassRepository
	userRepository   userDomain.UserRepository
	memberRepository memberDomain.MemberRepository
}

func New(classRepository classDomain.ClassRepository, userRepository userDomain.UserRepository, memberRepository memberDomain.MemberRepository) *MemberUseCases {
	return &MemberUseCases{
		classRepository:  classRepository,
		userRepository:   userRepository,
		memberRepository: memberRepository,
	}
}
