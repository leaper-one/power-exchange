package logic

import (
	"context"

	core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  更新用户
func (l *UpdateUserLogic) UpdateUser(in *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	user_new := &core.User{
		UUID: in.Uuid,
	}

	if in.UserName != "" {
		user_new.UserName = in.UserName
	}
	if in.Email != "" {
		user_new.Email = in.Email
	}
	if in.Phone != "" {
		user_new.Phone = in.Phone
	}

	err := l.svcCtx.DbEngin.Save(l.ctx, user_new)
	if err != nil {
		return &user.UpdateUserResponse{
			Code: 500,
			Msg:  "更新用户失败",
		}, nil
	}

	return &user.UpdateUserResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
