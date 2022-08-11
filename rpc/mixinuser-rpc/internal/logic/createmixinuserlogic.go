package logic

import (
	"context"

	"github.com/leaper-one/power-exchange/rpc/mixinuser-rpc/internal/svc"
	"github.com/leaper-one/power-exchange/rpc/mixinuser-rpc/mixinuser"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMixinUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMixinUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMixinUserLogic {
	return &CreateMixinUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  创建 MixinUser
func (l *CreateMixinUserLogic) CreateMixinUser(in *mixinuser.CreateMixinUserRequest) (*mixinuser.CreateMixinUserResponse, error) {
	// todo: add your logic here and delete this line

	return &mixinuser.CreateMixinUserResponse{}, nil
}
