package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/constants"
)

type GetClassCodeRequest struct {
	UserID  string
	ClassID string
}

func (useCases *ClassUseCases) GetClassCode(ctx context.Context, getClassCodeRequest *GetClassCodeRequest) (string, error) {
	if err := userDomain.IsIdValid(getClassCodeRequest.UserID); err != nil {
		return "", err
	}

	_, err := useCases.userRepository.Get(ctx, getClassCodeRequest.UserID)
	if err != nil {
		return "", err
	}

	if err := classDomain.IsIdValid(getClassCodeRequest.ClassID); err != nil {
		return "", err
	}

	memberRole, err := useCases.classRepository.GetMemberRoleByClassIDAndUserID(ctx, getClassCodeRequest.ClassID, getClassCodeRequest.UserID)
	if err != nil {
		return "", err
	}

	if memberRole != string(constants.RoleTeacher) {
		return "", classError.ErrUnauthorized
	}

	code, err := useCases.classRepository.GetClassCodeByClassID(ctx, getClassCodeRequest.ClassID)
	if err != nil {
		return "", err
	}

	return code, nil
}
