package logic

import (
	"context"

	"fizzbuzz/internal/svc"
	"fizzbuzz/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FizzbuzzLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFizzbuzzLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FizzbuzzLogic {
	return &FizzbuzzLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FizzbuzzLogic) Fizzbuzz(req *types.FizzbuzzRequest) (resp *types.FizzbuzzResponse, err error) {
	resp = &types.FizzbuzzResponse{}
	for _, msg := range l.svcCtx.Config.Messages {
		if req.Count > 0 && req.Count%msg.MultipleOf == 0 {
			resp.Message += msg.Message
		}
	}

	resp.Status = 200
	return
}
