package memberApplication

import (
	"context"

	classDomain "github.com/zchelalo/sa_class_management/internal/modules/class/domain"
	memberDomain "github.com/zchelalo/sa_class_management/internal/modules/member/domain"
	userDomain "github.com/zchelalo/sa_class_management/internal/modules/user/domain"
	"github.com/zchelalo/sa_class_management/pkg/bootstrap"
	"github.com/zchelalo/sa_class_management/pkg/meta"
)

type ListRequest struct {
	UserID  string
	ClassID string
	Page    int32
	Limit   int32
}

type ListResponse struct {
	Classes []*memberDomain.MemberEntity
	Meta    *meta.Meta
}

func (useCases *MemberUseCases) List(ctx context.Context, listRequest *ListRequest) (*ListResponse, error) {
	if err := userDomain.IsIdValid(listRequest.UserID); err != nil {
		return nil, err
	}

	if err := classDomain.IsIdValid(listRequest.ClassID); err != nil {
		return nil, err
	}

	if err := memberDomain.IsPageValid(listRequest.Page); err != nil {
		return nil, err
	}

	if err := memberDomain.IsLimitValid(listRequest.Limit); err != nil {
		return nil, err
	}

	_, err := useCases.userRepository.Get(ctx, listRequest.UserID)
	if err != nil {
		return nil, err
	}

	_, err = useCases.classRepository.GetByID(ctx, listRequest.ClassID)
	if err != nil {
		return nil, err
	}

	_, err = useCases.memberRepository.GetByUserIDAndClassID(ctx, listRequest.UserID, listRequest.ClassID)
	if err != nil {
		return nil, err
	}

	config := bootstrap.GetConfig()

	count, err := useCases.memberRepository.Count(ctx, listRequest.ClassID)
	if err != nil {
		return nil, err
	}

	meta, err := meta.New(listRequest.Page, listRequest.Limit, count, config.PaginatorLimitDefault)
	if err != nil {
		return nil, err
	}

	membersObtained, err := useCases.memberRepository.ListByClassID(ctx, listRequest.ClassID, meta.Offset(), meta.Limit())
	if err != nil {
		return nil, err
	}

	return &ListResponse{
		Classes: membersObtained,
		Meta:    meta,
	}, nil
}
