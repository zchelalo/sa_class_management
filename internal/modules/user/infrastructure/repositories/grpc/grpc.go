package userGRPC

import (
	"context"
	"errors"

	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	userError "github.com/zchelalo/sa_class_management/internal/modules/user/error"
	userProto "github.com/zchelalo/sa_class_management/pkg/proto/user"
	"google.golang.org/grpc/codes"
)

type GRPCRepository struct {
	client userProto.UserServiceClient
}

func NewGRPCRepository(client userProto.UserServiceClient) userDomain.UserRepository {
	return &GRPCRepository{
		client: client,
	}
}

func (r *GRPCRepository) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	user, err := r.client.GetUser(ctx, &userProto.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	errorObtained := user.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return nil, userError.ErrIdInvalid
		}
		if int32(codes.NotFound) == errorCode {
			return nil, userError.ErrUserNotFound
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	userObtained := user.GetUser()
	if userObtained == nil {
		return nil, userError.ErrUserNotFound
	}

	return &userDomain.UserEntity{
		ID:       userObtained.GetId(),
		Name:     userObtained.GetName(),
		Email:    userObtained.GetEmail(),
		Verified: userObtained.GetVerified(),
	}, nil
}
