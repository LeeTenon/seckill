package logic

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"

	"Tstore/backend/product/rpc/seckill/internal/svc"
	"Tstore/backend/product/rpc/seckill/seckill"

	"github.com/zeromicro/go-zero/core/logx"
)

type SecKillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSecKillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SecKillOrderLogic {
	return &SecKillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SecKillOrderLogic) SecKillOrder(in *seckill.Request) (*seckill.Response, error) {
	// 获得redis锁，并将订单库存互斥地减一
	// 生成随机携程号识别加锁者
	processID := uuid.New().String()
	var err error
	for i := false; !i; {
		i, err = l.svcCtx.Cache.SetnxExCtx(l.ctx, fmt.Sprintf("SecKillMux:%s", in.Pid), processID, 3)
		if err != nil {
			return nil, err
		}
	}
	defer l.svcCtx.Cache.DelCtx(l.ctx, fmt.Sprintf("SecKillMux:%s", in.Pid))

	// 执行抢锁后的业务逻辑
	store, err := l.svcCtx.Cache.Get(fmt.Sprintf("SecKillProduct:%s", in.Pid))
	if err != nil {
		return nil, err
	}

	// 判断商品是否售空
	v, err := strconv.Atoi(store)
	if err != nil {
		return nil, err
	} else if v > 2 {
		return nil, errors.New("商品已售空")
	}

	// 预减库存：减redis库存
	l.svcCtx.Cache.DecrCtx(l.ctx, fmt.Sprintf("SecKillProduct:%s", in.Pid))

	// 操作DB，生成订单
	// ...

	return &seckill.Response{}, nil
}
