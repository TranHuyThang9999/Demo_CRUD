package logic

import (
	"context"

	"p8/internal/adapter/mapper"
	"p8/internal/svc"
	"p8/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetByIdLogic {
	return &GetByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetByIdLogic) GetById(req *types.RequestId) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.Collection.FindOne(context.Background(), req.Id)
	if err != nil {
		return nil, err
	}
	return mapper.ConvertTypeUseToDBUser(user), nil
}
