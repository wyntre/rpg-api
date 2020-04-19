package models

func (ms *ModelSuite) CreateUser(email string, password string) *User {
	u := &User{
		Email:                email,
		Password:             password,
		PasswordConfirmation: password,
	}

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	return u
}

func (ms *ModelSuite) Test_User_Create() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")
	ms.NotNil(u.ID)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &User{
		Password: "password",
	}

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := ms.CreateUser("test@test.com", "test")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &User{
		Email:                "test@test.com",
		Password:             "test",
		PasswordConfirmation: "test",
	}

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}
