package memberGRPC

import (
	"context"

	classError "github.com/zchelalo/sa_class_management/internal/modules/class/error"
	memberApplication "github.com/zchelalo/sa_class_management/internal/modules/member/application"
	memberError "github.com/zchelalo/sa_class_management/internal/modules/member/error"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	memberProto "github.com/zchelalo/sa_class_management/pkg/proto/member"
	"github.com/zchelalo/sa_class_management/pkg/util"
	"google.golang.org/grpc/codes"
)

func (router *MemberRouter) ListMembers(ctx context.Context, req *memberProto.ListMembersRequest) (*memberProto.ListMembersResponse, error) {
	membersObtained, err := router.useCase.List(ctx, &memberApplication.ListRequest{
		UserID:  req.GetUserId(),
		ClassID: req.GetClassId(),
		Page:    req.GetPage(),
		Limit:   req.GetLimit(),
	})
	if err != nil {
		errorToReturn := &memberProto.MemberError{
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

		classErrorResponse := &memberProto.ListMembersResponse{
			Result: &memberProto.ListMembersResponse_Error{
				Error: errorToReturn,
			},
		}

		return classErrorResponse, nil
	}

	members := make([]*memberProto.MemberData, 0)
	for _, memberObtained := range membersObtained.Members {
		member := &memberProto.MemberData{
			Id: memberObtained.ID,
			User: &memberProto.MemberUserData{
				Id:       memberObtained.User.ID,
				Name:     memberObtained.User.Name,
				Email:    memberObtained.User.Email,
				Verified: memberObtained.User.Verified,
			},
			Role: &memberProto.MemberRoleData{
				Id:  memberObtained.Role.ID,
				Key: memberObtained.Role.Key,
			},
		}
		members = append(members, member)
	}

	response := &memberProto.ListMembersResponse{
		Result: &memberProto.ListMembersResponse_Data{
			Data: &memberProto.MembersWithMeta{
				Members: members,
				Meta: &memberProto.MemberMeta{
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
