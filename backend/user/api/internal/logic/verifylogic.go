package logic

import (
	"context"

	"Tstore/backend/user/api/internal/svc"
	"Tstore/backend/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLogic {
	return &VerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLogic) Verify(req *types.VerifyReq) (resp *types.VerifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
