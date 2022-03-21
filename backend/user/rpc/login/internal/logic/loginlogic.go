package logic

import (
	"Tstore/dao"
	"Tstore/utils"
	"context"
	"errors"
	"time"

	"Tstore/backend/user/rpc/login/internal/svc"
	"Tstore/backend/user/rpc/login/login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *login.LoginReq) (*login.LoginResp, error) {
	// 校验密码
	result, err := dao.UserDao.FondByName(l.svcCtx.DbEngin, in.UserName)
	if err != nil {
		if err == dao.ErrorNoData {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}
	if in.Password != result.Password {
		return nil, errors.New("用户名或密码错误")
	}

	// 获取Token
	token, err := utils.GetJwtToken(l.svcCtx.Config.AccessSecret, time.Now().Unix(), l.svcCtx.Config.AccessExpire, int64(result.ID))

	// 返回消息
	resp := &login.LoginResp{
		Id:     int64(result.ID),
		Name:   result.UserName,
		Avatar: result.Avatar,
		Token:  token,
	}
	return resp, nil
}
