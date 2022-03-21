package logic

import (
	"Tstore/backend/user/rpc/login/loginclient"
	"context"

	"Tstore/backend/user/api/internal/svc"
	"Tstore/backend/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResp, error) {
	result, err := l.svcCtx.Login.Login(l.ctx, &loginclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp := &types.LoginResp{
		ID:     int(result.Id),
		Name:   result.Name,
		Avatar: result.Avatar,
		Token:  result.Token,
	}

	return resp, nil
}
