package logic

import (
	"context"

	"Tstore/backend/product/api/internal/svc"
	"Tstore/backend/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillLogic {
	return &SeckillLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillLogic) Seckill(req *types.SeckillReq) (resp *types.SeckillResp, err error) {
	// todo: add your logic here and delete this line

	return
}
