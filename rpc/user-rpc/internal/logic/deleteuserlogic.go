package logic

import (
	"context"

	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  删除用户
func (l *DeleteUserLogic) DeleteUser(in *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	err := l.svcCtx.DbEngin.DeleteByUUID(l.ctx, in.Uuid)
	if err != nil {
		return &user.DeleteUserResponse{
			Code: 500,
			Msg:  "删除用户失败",
		}, err
	}

	return &user.DeleteUserResponse{}, nil
}
