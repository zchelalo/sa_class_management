package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/constants"
)

type JoinRequest struct {
	UserID string
	Code   string
}

func (useCases *ClassUseCases) Join(ctx context.Context, joinRequest *JoinRequest) (*classDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(joinRequest.UserID); err != nil {
		return nil, err
	}

	_, err := useCases.userRepository.Get(ctx, joinRequest.UserID)
	if err != nil {
		return nil, err
	}

	classObtained, err := useCases.classRepository.GetClassByCode(ctx, joinRequest.Code)
	if err != nil {
		return nil, err
	}

	_, err = useCases.memberRepository.GetByUserIDAndClassID(ctx, joinRequest.UserID, classObtained.ID)
	if err != memberError.ErrMemberNotFound {
		return nil, classError.ErrAlreadyJoined
	}

	studentRole, err := useCases.memberRepository.GetRoleIDByKey(ctx, string(constants.RoleStudent))
	if err != nil {
		return nil, err
	}

	newMember, err := memberDomain.New(joinRequest.UserID, studentRole)
	if err != nil {
		return nil, err
	}

	err = useCases.classRepository.Join(ctx, newMember, classObtained.ID)
	if err != nil {
		return nil, err
	}

	return classObtained, nil
}
