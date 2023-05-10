package logic

import (
	"context"

	"p8/internal/svc"
	"p8/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteByIdLogic {
	return &DeleteByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteByIdLogic) DeleteById(req *types.User) (resp *types.User, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.Collection.Delete(l.ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.User{
		Id:       req.Id,
		UserName: req.UserName,
		Address:  req.Address,
		Gender:   req.Gender,
	}, nil
}
