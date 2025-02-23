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

func (router *ClassRouter) JoinClass(ctx context.Context, req *proto.JoinClassRequest) (*proto.JoinClassResponse, error) {
	classObtained, err := router.useCase.Join(ctx, &classApplication.JoinRequest{
		UserID: req.GetUserId(),
		Code:   req.GetCode(),
	})
	if err != nil {
		errorToReturn := &proto.Error{
			Code:    int32(codes.Internal),
			Message: "Internal server error",
		}

		badRequestErrors := []error{
			userError.ErrIdRequired,
			userError.ErrIdInvalid,

			classError.ErrCodeRequired,
			classError.ErrCodeInvalid,
		}

		if util.IsErrorType(err, badRequestErrors) {
			errorToReturn.Code = int32(codes.InvalidArgument)
			errorToReturn.Message = err.Error()
		}

		if err == userError.ErrUserNotFound || err == classError.ErrClassNotFound {
			errorToReturn.Code = int32(codes.NotFound)
			errorToReturn.Message = err.Error()
		}

		if err == classError.ErrAlreadyJoined {
			errorToReturn.Code = int32(codes.AlreadyExists)
			errorToReturn.Message = err.Error()
		}

		classErrorResponse := &proto.JoinClassResponse{
			Result: &proto.JoinClassResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	response := &proto.JoinClassResponse{
		Result: &proto.JoinClassResponse_Class{
			Class: &proto.ClassData{
				Id:      classObtained.ID,
				Name:    classObtained.Name,
				Subject: classObtained.Subject,
				Grade:   classObtained.Grade,
			},
		},
	}

	return response, nil
}
