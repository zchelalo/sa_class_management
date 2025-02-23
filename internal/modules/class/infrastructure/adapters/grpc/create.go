package classGRPC

import (
	"context"

	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	"github.com/zchelalo/sa_class_management/pkg/proto"
	"github.com/zchelalo/sa_class_management/pkg/util"
	"google.golang.org/grpc/codes"
)

func (router *ClassRouter) CreateClass(ctx context.Context, req *proto.CreateClassRequest) (*proto.CreateClassResponse, error) {
	classCreated, err := router.useCase.Create(ctx, &classApplication.CreateRequest{
		UserID:  req.GetUserId(),
		Name:    req.GetName(),
		Subject: req.GetSubject(),
		Grade:   req.GetGrade(),
	})
	if err != nil {
		errorToReturn := &proto.Error{
			Code:    int32(codes.Internal),
			Message: "Internal server error",
		}

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
			errorToReturn.Code = int32(codes.InvalidArgument)
			errorToReturn.Message = err.Error()
		}

		if err == userError.ErrUserNotFound {
			errorToReturn.Code = int32(codes.NotFound)
			errorToReturn.Message = err.Error()
		}

		classErrorResponse := &proto.CreateClassResponse{
			Result: &proto.CreateClassResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	response := &proto.CreateClassResponse{
		Result: &proto.CreateClassResponse_Class{
			Class: &proto.ClassData{
				Id:      classCreated.ID,
				Name:    classCreated.Name,
				Subject: classCreated.Subject,
				Grade:   classCreated.Grade,
			},
		},
	}

	return response, nil
}
