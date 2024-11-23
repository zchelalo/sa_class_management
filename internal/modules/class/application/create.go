package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/constants"
)

type CreateRequest struct {
	UserID  string
	Name    string
	Subject string
	Grade   string
}

func (useCases *ClassUseCases) Create(ctx context.Context, createRequest *CreateRequest) (*classDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(createRequest.UserID); err != nil {
		return nil, err
	}

	_, err := useCases.userRepository.Get(ctx, createRequest.UserID)
	if err != nil {
		return nil, err
	}

	newClass, err := classDomain.New(createRequest.Name, createRequest.Subject, createRequest.Grade)
	if err != nil {
		return nil, err
	}

	teacherRole, err := useCases.memberRepository.GetRoleIDByKey(ctx, string(constants.RoleTeacher))
	if err != nil {
		return nil, err
	}

	newMember, err := memberDomain.New(createRequest.UserID, teacherRole)
	if err != nil {
		return nil, err
	}

	classCreated, err := useCases.classRepository.Create(ctx, newMember, newClass)
	if err != nil {
		return nil, err
	}

	return classCreated, nil
}
