package actions

import (
	"net/http"

	"github.com/wyntre/rpg_api/models"
)

// test no token provided
func (as *ActionSuite) Test_HomeHandler() {
	res := as.JSON("/").Get()
	as.Equal(http.StatusUnauthorized, res.Code)
}

// test valid auth token
func (as *ActionSuite) Test_HomeHandler_Authorized() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	// auth'd request
	req := as.JSON("/")
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusOK, res.Code)
}
