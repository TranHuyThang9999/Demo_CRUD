package mapper

import (
	"database/sql"
	"p8/internal/adapter/database"
	"p8/internal/types"
)

func ConvertToUser(u *types.User) *database.User {
	return &database.User{
		Id:       u.Id,
		UserName: sql.NullString{String: u.UserName, Valid: true},
		Password: sql.NullString{String: u.Password, Valid: true},
		Age:      sql.NullInt64{Int64: u.Age, Valid: true},
		Address:  sql.NullString{String: u.Address, Valid: true},
		Gender:   sql.NullString{String: u.Gender, Valid: true},
	}
}

func ConvertUserDbToUserType(u database.User) types.User {
	return types.User{
		Id:       u.Id,
		UserName: u.UserName.String,
		Password: u.Password.String,
		Age:      u.Age.Int64,
		Address:  u.Address.String,
		Gender:   u.Gender.String,
	}
}
func ConvertUserDbToUserType1(u *database.User) *types.User {
	return &types.User{
		Id:       u.Id,
		UserName: u.UserName.String,
		Password: u.Password.String,
		Age:      u.Age.Int64,
		Address:  u.Address.String,
		Gender:   u.Gender.String,
	}
}
func ConvertListUserDbToListUserType(users []database.User) []types.User {
	result := make([]types.User, len(users))
	for i, u := range users {
		result[i] = ConvertUserDbToUserType(u)
	}
	return result
}

// func ConvertToUserLogin(u *types.LoginReq) *databasecustomer.User {
// 	return &databasecustomer.User{
// 		UserName: sql.NullString{String: u.UserName, Valid: true},
// 		Password: sql.NullString{String: u.Password, Valid: true},
// 	}
// }
