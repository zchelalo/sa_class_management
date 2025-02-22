package classApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/meta"
)

type ListRequest struct {
	UserID string
	Page   int32
	Limit  int32
}

type ListResponse struct {
	Classes []*classDomain.ClassEntity
	Meta    *meta.Meta
}

func (useCases *ClassUseCases) List(ctx context.Context, listRequest *ListRequest) (*ListResponse, error) {
	if err := userDomain.IsIdValid(listRequest.UserID); err != nil {
		return nil, err
	}

	if err := classDomain.IsPageValid(listRequest.Page); err != nil {
		return nil, err
	}

	if err := classDomain.IsLimitValid(listRequest.Limit); err != nil {
		return nil, err
	}

	_, err := useCases.userRepository.Get(ctx, listRequest.UserID)
	if err != nil {
		return nil, err
	}

	config := bootstrap.GetConfig()

	count, err := useCases.classRepository.Count(ctx, listRequest.UserID)
	if err != nil {
		return nil, err
	}

	meta, err := meta.New(listRequest.Page, listRequest.Limit, count, config.PaginatorLimitDefault)
	if err != nil {
		return nil, err
	}

	classesObtained, err := useCases.classRepository.List(ctx, listRequest.UserID, meta.Offset(), meta.Limit())
	if err != nil {
		return nil, err
	}

	return &ListResponse{
		Classes: classesObtained,
		Meta:    meta,
	}, nil
}
