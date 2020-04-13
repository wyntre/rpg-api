package actions

import (
  "net/http"
)

type CreateUserRequest struct {
  Email                 string `json:"email"`
  Password              string `json:"password"`
  PasswordConfirmation 	string `json:"password_confirmation"`
}

type CreateUserFailureResponse struct {
  Code  string `json:"code"`
  Error string `json:"error"`
  Trace string `json:"trace"`
}

func (as *ActionSuite) Test_Users_Create() {
  u := &CreateUserRequest{
    Email:                "test@test.com",
    Password:             "test",
    PasswordConfirmation: "test",
  }

  t := &AuthTokenResponse{}

  res := as.JSON("/v1/users").Post(u)
  res.Bind(t)
  as.Equal(http.StatusCreated, res.Code)
  as.NotNil(t.Token)
}

func (as *ActionSuite) Test_Users_Create_Bad_Email() {
  u := &CreateUserRequest{
    Email: "test",
    Password: "test",
    PasswordConfirmation: "test",
  }

  e := &CreateUserFailureResponse{}

  res := as.JSON("/v1/users").Post(u)
  res.Bind(e)
  as.Equal(http.StatusConflict, res.Code)
  as.Equal("Incorrect email format", e.Error)
}

func (as *ActionSuite) Test_Users_Create_Bad_Password() {
  u := &CreateUserRequest{
    Email: "test@test.com",
    Password: "test",
    PasswordConfirmation: "test2",
  }

  e := &CreateUserFailureResponse{}

  res := as.JSON("/v1/users").Post(u)
  res.Bind(e)
  as.Equal(http.StatusConflict, res.Code)
  as.Equal("Password does not match confirmation", e.Error)
}
