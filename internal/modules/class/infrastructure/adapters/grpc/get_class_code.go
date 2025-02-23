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

func (router *ClassRouter) GetClassCode(ctx context.Context, req *proto.GetClassCodeRequest) (*proto.GetClassCodeResponse, error) {
	classCodeObtained, err := router.useCase.GetClassCode(ctx, &classApplication.GetClassCodeRequest{
		UserID:  req.GetUserId(),
		ClassID: req.GetClassId(),
	})
	if err != nil {
		errorToReturn := &proto.Error{
			Code:    int32(codes.Internal),
			Message: "Internal server error",
		}

		badRequestErrors := []error{
			userError.ErrIdRequired,
			userError.ErrIdInvalid,

			classError.ErrIdInvalid,
			classError.ErrIdRequired,
		}

		if util.IsErrorType(err, badRequestErrors) {
			errorToReturn.Code = int32(codes.InvalidArgument)
			errorToReturn.Message = err.Error()
		}

		if err == userError.ErrUserNotFound || err == classError.ErrClassNotFound {
			errorToReturn.Code = int32(codes.NotFound)
			errorToReturn.Message = err.Error()
		}

		if err == classError.ErrUnauthorized {
			errorToReturn.Code = int32(codes.PermissionDenied)
			errorToReturn.Message = err.Error()
		}

		classErrorResponse := &proto.GetClassCodeResponse{
			Result: &proto.GetClassCodeResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	response := &proto.GetClassCodeResponse{
		Result: &proto.GetClassCodeResponse_Code{
			Code: classCodeObtained,
		},
	}

	return response, nil
}
