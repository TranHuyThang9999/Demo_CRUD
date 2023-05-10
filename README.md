func (l *CreateUserLogic) CreateUser(req *types.UserRequest) (resp *types.UserReponse, err error) {

	now := time.Now()
	//	req. = now.Unix()
	hash, _ := bootrap.HashPassword(req.Password)
	data := types.User{
		Id:       now.Unix(),
		UserName: req.UserName,
		Password: hash,
		Age:      req.Age,
		Address:  req.Address,
		Gender:   req.Gender,
	}
	_, err := l.svcCtx.Collection.Insert(context.Background(), mapper.ConvertToUser(&data))
	if err != nil {
		return nil, err
	}
	return resp, nil
}
@handler CreateUser // TODO: set handler name and delete this comment
	post /users/create(UserRequest) returns (UserReponse)

	//md5
	//dh key exchange algorithm


	func (m *defaultUserModel) Login(ctx context.Context, user *types.LoginReq) (*types.LoginReply, error) {
	var userModel User
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `user_name` = '%s' LIMIT 1", userRows, m.table, user.UserName)
	err := m.conn.QueryRowCtx(ctx, &userModel, query)

	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userModel.Password), []byte(user.Password)); err != nil {
		return nil, nil
	}

	return nil, nil
}
