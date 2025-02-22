package classGRPC

import (
	"context"

	classApplication "github.com/zchelalo/sa_class_management/internal/modules/class/application"
	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	classProto "github.com/zchelalo/sa_class_management/pkg/proto/class"
	"github.com/zchelalo/sa_class_management/pkg/util"
	"google.golang.org/grpc/codes"
)

func (router *ClassRouter) ListClasses(ctx context.Context, req *classProto.ListClassesRequest) (*classProto.ListClassesResponse, error) {
	classesObtained, err := router.useCase.List(ctx, &classApplication.ListRequest{
		UserID: req.GetUserId(),
		Page:   req.GetPage(),
		Limit:  req.GetLimit(),
	})
	if err != nil {
		errorToReturn := &classProto.ClassError{
			Code:    int32(codes.Internal),
			Message: "Internal server error",
		}

		badRequestErrors := []error{
			userError.ErrIdRequired,
			userError.ErrIdInvalid,

			classError.ErrPageInvalid,
			classError.ErrLimitInvalid,
		}

		if util.IsErrorType(err, badRequestErrors) {
			errorToReturn.Code = int32(codes.InvalidArgument)
			errorToReturn.Message = err.Error()
		}

		if err == userError.ErrUserNotFound || err == classError.ErrClassesNotFound {
			errorToReturn.Code = int32(codes.NotFound)
			errorToReturn.Message = err.Error()
		}

		classErrorResponse := &classProto.ListClassesResponse{
			Result: &classProto.ListClassesResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	classes := make([]*classProto.ClassData, 0)
	for _, classObtained := range classesObtained.Classes {
		class := &classProto.ClassData{
			Id:      classObtained.ID,
			Name:    classObtained.Name,
			Subject: classObtained.Subject,
			Grade:   classObtained.Grade,
		}
		classes = append(classes, class)
	}

	response := &classProto.ListClassesResponse{
		Result: &classProto.ListClassesResponse_Data{
			Data: &classProto.ClassesWithMeta{
				Classes: classes,
				Meta: &classProto.ClassMeta{
					Page:       classesObtained.Meta.Page,
					PerPage:    classesObtained.Meta.PerPage,
					Count:      classesObtained.Meta.PageCount,
					TotalCount: classesObtained.Meta.TotalCount,
				},
			},
		},
	}

	return response, nil
}
