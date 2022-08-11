package logic

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/leaper-one/pkg/util"
	core "github.com/leaper-one/power-exchange/core/user"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/internal/svc"
	"github.com/leaper-one/power-exchange/rpc/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  创建用户
func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	// UUID
	uuid, _ := uuid.NewV1()
	// 密码加密
	password, err := util.HashPassword(in.Password)
	if err != nil {
		return &user.CreateUserResponse{
			Code: 400,
			Msg:  "密码加密失败",
		}, nil
	}

	new_user := core.User{
		UUID:      uuid.String(),
		Password:  password,
		UserName:  in.UserName,
		Email:     in.Email,
		Phone:     in.Phone,
		AvatarUrl: in.Avatar,
		Type:      in.Type,
	}

	// 存入数据库
	err = l.svcCtx.DbEngin.Save(l.ctx, &new_user)
	if err != nil {
		return &user.CreateUserResponse{
			Code: 500,
			Msg:  "数据库创建用户失败",
		}, err
	}
	return &user.CreateUserResponse{
		Code: 200,
		Msg:  "创建用户成功",
	}, err
}
