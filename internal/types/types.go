// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

type UserReponse struct {
	Message string `json:"message"`
}

type UserRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
}

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginReply struct {
	Id           int64  `json:"id"`
	UserName     string `json:"user_name"`
	Gender       string `json:"gender"`
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
	RefreshAfter int64  `json:"refreshAfter"`
}

type RequestId struct {
	Id int64 `json:"id"`
}

type SearchReq struct {
	Name string `form:"name"`
}

type SearchReply struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type ReponseListUser struct {
	User []User `json:"user"`
}
