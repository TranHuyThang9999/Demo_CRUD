package database

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		GetByUserName(ctx context.Context, userName string) (*User, error)
		GetAllUser(ctx context.Context) ([]User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}
func (l *defaultUserModel) GetByUserName(ctx context.Context, userName string) (*User, error) {
	query := fmt.Sprintf("select %s from %s where `user_name` = ? limit 1", userRows, l.table)
	var resp User
	err := l.conn.QueryRowCtx(ctx, &resp, query, userName)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

//	func (l *defaultUserModel) GetAllUser(ctx context.Context) ([]User, error) {
//		var users = make([]User, 0)
//		query := fmt.Sprintf("select %s from %s", userRows, l.table)
//		err := l.conn.QueryRowCtx(ctx, &users, query)
//		switch err {
//		case nil:
//			return users, nil
//		case sqlc.ErrNotFound:
//			return nil, ErrNotFound
//		default:
//			return nil, err
//		}
//	}
func (l *defaultUserModel) GetAllUser(ctx context.Context) ([]User, error) {
	var users []User
	query := fmt.Sprintf("SELECT * FROM %s", l.table)
	err := l.conn.QueryRowsCtx(ctx, &users, query)
	switch err {
	case nil:
		return users, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
