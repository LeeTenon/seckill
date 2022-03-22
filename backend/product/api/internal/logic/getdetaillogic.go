package logic

import (
	"context"

	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailLogic {
	return &GetDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDetailLogic) GetDetail(req *types.GetDetailReq) (resp *types.GetDetailResp, err error) {
	// todo: add your logic here and delete this line

	return
}
