package classGRPC

import (
	"context"

	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	classProto "github.com/zchelalo/sa_class_management/pkg/proto/class"
	"github.com/zchelalo/sa_class_management/pkg/util"
)

func (router *ClassRouter) CreateClass(ctx context.Context, req *classProto.CreateClassRequest) (*classProto.CreateClassResponse, error) {
	classCreated, err := router.useCase.Create(ctx, &classApplication.CreateRequest{
		UserID:  req.GetUserId(),
		Name:    req.GetName(),
		Subject: req.GetSubject(),
		Grade:   req.GetGrade(),
	})
	if err != nil {
		classErrorResponse := &classProto.ClassError{}

		badRequestErrors := []error{
			classError.ErrNameRequired,
			classError.ErrNameTooShort,
			classError.ErrSubjectRequired,
			classError.ErrSubjectTooShort,
			classError.ErrGradeRequired,
			classError.ErrGradeTooShort,

			userError.ErrIdRequired,
			userError.ErrIdInvalid,
		}

		if util.IsErrorType(err, badRequestErrors) {

		}

		return nil, err
	}

	response := &classProto.CreateClassResponse{
		Result: &classProto.CreateClassResponse_Class{
			Class: &classProto.ClassData{
				Id:      classCreated.ID,
				Name:    classCreated.Name,
				Subject: classCreated.Subject,
				Grade:   classCreated.Grade,
			},
		},
	}

	return response, nil
}
