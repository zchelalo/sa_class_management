package memberGRPC

import (
	"context"

	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	memberApplication "github.com/zchelalo/sa_class_management/internal/modules/member/application"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	"github.com/zchelalo/sa_class_management/pkg/proto"
	"github.com/zchelalo/sa_class_management/pkg/util"
	"google.golang.org/grpc/codes"
)

func (router *MemberRouter) ListMembers(ctx context.Context, req *proto.ListMembersRequest) (*proto.ListMembersResponse, error) {
	membersObtained, err := router.useCase.List(ctx, &memberApplication.ListRequest{
		UserID:  req.GetUserId(),
		ClassID: req.GetClassId(),
		Page:    req.GetPage(),
		Limit:   req.GetLimit(),
	})
	if err != nil {
		errorToReturn := &proto.Error{
			Code:    int32(codes.Internal),
			Message: "Internal server error",
		}

		badRequestErrors := []error{
			userError.ErrIdRequired,
			userError.ErrIdInvalid,

			classError.ErrIdRequired,
			classError.ErrIdInvalid,

			memberError.ErrPageInvalid,
			memberError.ErrLimitInvalid,
		}

		if util.IsErrorType(err, badRequestErrors) {
			errorToReturn.Code = int32(codes.InvalidArgument)
			errorToReturn.Message = err.Error()
		}

		notFoundErrors := []error{
			userError.ErrUserNotFound,
			classError.ErrClassNotFound,
			memberError.ErrMemberNotFound,
			memberError.ErrMembersNotFound,
		}

		if util.IsErrorType(err, notFoundErrors) {
			errorToReturn.Code = int32(codes.NotFound)
			errorToReturn.Message = err.Error()
		}

		classErrorResponse := &proto.ListMembersResponse{
			Result: &proto.ListMembersResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	members := make([]*proto.MemberData, 0)
	for _, memberObtained := range membersObtained.Members {
		member := &proto.MemberData{
			Id: memberObtained.ID,
			User: &proto.UserData{
				Id:       memberObtained.User.ID,
				Name:     memberObtained.User.Name,
				Email:    memberObtained.User.Email,
				Verified: memberObtained.User.Verified,
			},
			Role: &proto.MemberRoleData{
				Id:  memberObtained.Role.ID,
				Key: memberObtained.Role.Key,
			},
		}
		members = append(members, member)
	}

	response := &proto.ListMembersResponse{
		Result: &proto.ListMembersResponse_Data{
			Data: &proto.MembersWithMeta{
				Members: members,
				Meta: &proto.Meta{
					Page:       membersObtained.Meta.Page,
					PerPage:    membersObtained.Meta.PerPage,
					Count:      membersObtained.Meta.PageCount,
					TotalCount: membersObtained.Meta.TotalCount,
				},
			},
		},
	}

	return response, nil
}
