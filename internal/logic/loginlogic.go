package logic

import (
	"context"
	"time"

	"p8/internal/adapter/database"
	"p8/internal/svc"
	"p8/internal/types"

	"github.com/golang-jwt/jwt"
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
	//var token string

	user, err = l.svcCtx.Collection.GetByUserName(l.ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	result := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(req.Password))
	if result != nil {
		return nil, result
	}

	// token, err = middware.GenerateJWT(req.UserName, req.Password)
	// if err != nil {
	// 	return nil, err
	// }
	// l.Info(token)

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, user.Id)
	if err != nil {
		return nil, err
	}

	return &types.LoginReply{
		Id:           user.Id,
		UserName:     user.UserName.String,
		Gender:       user.Gender.String,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil

}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
