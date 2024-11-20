package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
)

func (useCases *ClassUseCases) Create(ctx context.Context, userID, name, subject, grade string) (*classDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(userID); err != nil {
		return nil, err
	}

	_, err := useCases.UserRepository.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	newClass, err := classDomain.New(name, subject, grade)
	if err != nil {
		return nil, err
	}

	classCreated, err := useCases.classRepository.Create(ctx, userID, newClass)
	if err != nil {
		return nil, err
	}

	return classCreated, nil
}
