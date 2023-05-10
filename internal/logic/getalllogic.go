package logic

import (
	"context"

	"p8/internal/adapter/mapper"
	"p8/internal/svc"
	"p8/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllLogic) GetAll() (resp *types.ReponseListUser, err error) {
	// todo: add your logic here and delete this line
	user, err := l.svcCtx.Collection.GetAllUser(l.ctx)
	if err != nil {
		return nil, err
	}
	//	logx.Info(user)
	return &types.ReponseListUser{
		User: mapper.ConvertListUserDbToListUserType(user),
	}, nil
}
