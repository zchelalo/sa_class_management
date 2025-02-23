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

func (router *ClassRouter) ListClasses(ctx context.Context, req *proto.ListClassesRequest) (*proto.ListClassesResponse, error) {
	classesObtained, err := router.useCase.List(ctx, &classApplication.ListRequest{
		UserID: req.GetUserId(),
		Page:   req.GetPage(),
		Limit:  req.GetLimit(),
	})
	if err != nil {
		errorToReturn := &proto.Error{
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

		classErrorResponse := &proto.ListClassesResponse{
			Result: &proto.ListClassesResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	classes := make([]*proto.ClassData, 0)
	for _, classObtained := range classesObtained.Classes {
		class := &proto.ClassData{
			Id:      classObtained.ID,
			Name:    classObtained.Name,
			Subject: classObtained.Subject,
			Grade:   classObtained.Grade,
		}
		classes = append(classes, class)
	}

	response := &proto.ListClassesResponse{
		Result: &proto.ListClassesResponse_Data{
			Data: &proto.ClassesWithMeta{
				Classes: classes,
				Meta: &proto.Meta{
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
