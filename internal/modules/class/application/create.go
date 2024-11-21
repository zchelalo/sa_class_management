package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/constants"
)

func (useCases *ClassUseCases) Create(ctx context.Context, userID, name, subject, grade string) (*classDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(userID); err != nil {
		return nil, err
	}

	_, err := useCases.userRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	newClass, err := classDomain.New(name, subject, grade)
	if err != nil {
		return nil, err
	}

	teacherRole, err := useCases.memberRepository.GetRoleIDByKey(ctx, string(constants.RoleTeacher))
	if err != nil {
		return nil, err
	}

	newMember, err := memberDomain.New(userID, teacherRole)
	if err != nil {
		return nil, err
	}

	classCreated, err := useCases.classRepository.Create(ctx, newMember, newClass)
	if err != nil {
		return nil, err
	}

	return classCreated, nil
}
