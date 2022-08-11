package logic

import (
	"context"

	// core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserLogic {
	return &QueryUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  查询用户
func (l *QueryUserLogic) QueryUser(in *user.QueryUserRequest) (*user.QueryUserResponse, error) {
	// 查询数据库
	user_res, err := l.svcCtx.DbEngin.FindByUUID(l.ctx, in.Uuid)
	if user_res == nil && err == nil {
		return &user.QueryUserResponse{
			Code: 400,
			Msg:  "该 UUID 不存在",
		}, err
	} else if err != nil {
		return &user.QueryUserResponse{
			Code: 500,
			Msg:  "查询用户失败",
		}, err
	}

	return &user.QueryUserResponse{
		Code:   200,
		Msg:    "success",
		Name:   user_res.UserName,
		Email:  user_res.Email,
		Phone:  user_res.Phone,
		Avatar: user_res.AvatarUrl,
	}, nil
}
