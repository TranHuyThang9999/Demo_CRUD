package logic

import (
	"context"

	"p8/internal/adapter/database"
	"p8/internal/svc"
	"p8/internal/types"
	"p8/middware"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	var user *database.User
	var token string

	user, err = l.svcCtx.Collection.GetByUserName(l.ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	result := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password))
	if result != nil {
		return nil, result
	}

	token, err = middware.GenerateJWT(req.UserName, req.Password)
	if err != nil {
		return nil, err
	}
	l.Info(token)
	return &types.LoginReply{
		Code:        "",
		AccessToken: token,
	}, nil

}
