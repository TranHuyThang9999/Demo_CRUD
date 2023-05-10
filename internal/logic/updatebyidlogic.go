package logic

import (
	"context"

	"p8/internal/adapter/mapper"
	"p8/internal/svc"
	"p8/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateByIdLogic {
	return &UpdateByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateByIdLogic) UpdateById(req *types.User) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.Collection.Update(l.ctx, mapper.ConvertToUser(req))
	if err != nil {
		return nil, nil
	}
	return &types.User{
		Id:       req.Id,
		UserName: req.UserName,
		Address:  req.Address,
	}, nil
}
