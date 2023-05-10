package logic

import (
	"context"
	"time"

	"p8/internal/adapter/mapper"
	"p8/internal/svc"
	"p8/internal/types"

	"p8/internal/bootstrap"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.UserRequest) (resp bool, err error) {
	// todo: add your logic here and delete this line
	now := time.Now()
	//	req. = now.Unix()
	hash, _ := bootstrap.HashPassword(req.Password)
	data := types.User{
		Id:       now.Unix(),
		UserName: req.UserName,
		Password: hash,
		Age:      req.Age,
		Address:  req.Address,
		Gender:   req.Gender,
	}
	_, err = l.svcCtx.Collection.Insert(context.Background(), mapper.ConvertToUser(&data))
	if err != nil {
		return false, err
	}
	return true, nil
}
