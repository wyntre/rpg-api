package models

func (ms *ModelSuite) Test_Revoke_Token_Create() {
	count, err := ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(0, count)

	token := &Revokedtoken{
		Token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU2MDAyOTAsImlkIjoiNWY0MjlhZWEtZmFlMy00ZjU5LTlkNGUtMmY0ZTY3YTExZTU2In0.Dm1oKJTpcLi4ASlbK9YE-stCV6PcUNwerntSYn5J2VNArKdw_QRqRXyI5ijsqZYKIwj3Dk18y9knCPpYxrzyjsveBVPvKuZmvSuXPJ-v3kiPYSRv-HrRrKS0FzODMrNWM90exKP_KDQUIBrmRuvMMB2Lp9AqEvt3rzjbhetrdS-4Rb4ay1X9VzI35LBffQQxyYLjzrZewccKaBRqpmwoFZUsM8ht2b6nKBjOfnEhR6w5FHgqhZmNy6C8wt7zdOYYKmtJ3yzo3V6mcEFn1_DwnrjQ_AZ6igHlaFYbkCaIoISsmVZQ8BmkgBneP5RJo34ORIZXrZDlzF8rmcZtwfnppQ",
	}

	verrs, err := ms.DB.ValidateAndCreate(token)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Revoke_Token_ValidationErrors() {
	count, err := ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(0, count)

	token := &Revokedtoken{}

	verrs, err := ms.DB.ValidateAndCreate(token)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Revoke_Token_TokenExists() {
	count, err := ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(0, count)

	token := &Revokedtoken{
		Token: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU2MDAyOTAsImlkIjoiNWY0MjlhZWEtZmFlMy00ZjU5LTlkNGUtMmY0ZTY3YTExZTU2In0.Dm1oKJTpcLi4ASlbK9YE-stCV6PcUNwerntSYn5J2VNArKdw_QRqRXyI5ijsqZYKIwj3Dk18y9knCPpYxrzyjsveBVPvKuZmvSuXPJ-v3kiPYSRv-HrRrKS0FzODMrNWM90exKP_KDQUIBrmRuvMMB2Lp9AqEvt3rzjbhetrdS-4Rb4ay1X9VzI35LBffQQxyYLjzrZewccKaBRqpmwoFZUsM8ht2b6nKBjOfnEhR6w5FHgqhZmNy6C8wt7zdOYYKmtJ3yzo3V6mcEFn1_DwnrjQ_AZ6igHlaFYbkCaIoISsmVZQ8BmkgBneP5RJo34ORIZXrZDlzF8rmcZtwfaaaa",
	}

	verrs, err := ms.DB.ValidateAndCreate(token)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	count, err = ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(1, count)

	verrs, err = ms.DB.ValidateAndCreate(token)
	ms.Error(err)

	count, err = ms.DB.Count("revokedtokens")
	ms.NoError(err)
	ms.Equal(1, count)
}
