package actions

import (
	"net/http"

	"github.com/wyntre/rpg_api/models"
)

type CreateCampaignRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CampaignsListResponse struct {
	Campaigns models.Campaigns `json:"campaigns"`
}

func (as *ActionSuite) CreateCampaign(name string, description string, token string) *models.Campaign {
	// create campaign
	campaignRequest := &CreateCampaignRequest{
		Name:        name,
		Description: description,
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)

	return campaign
}

func (as *ActionSuite) Test_Campaigns_Create() {
	token := as.CreateUser("test@test.com", "test")

	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)
}

func (as *ActionSuite) Test_Campaigns_Create_Fail() {
	token := as.CreateUser("test@test.com", "test")

	campaignRequest := &CreateCampaignRequest{
		Name: "Test Campaign",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Create_No_Token() {
	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	res := req.Post(campaignRequest)
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Show() {
	token := as.CreateUser("test@test.com", "test")

	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)

	req = as.JSON("/v1/campaigns/show/" + campaign.ID.String())
	req.Headers["Authorization"] = token
	res = req.Get()
	as.Equal(http.StatusOK, res.Code)

	testCampaign := &models.Campaign{}
	res.Bind(testCampaign)
	as.Equal(campaign.ID, testCampaign.ID)
	as.Equal(campaign.Name, testCampaign.Name)
}

func (as *ActionSuite) Test_Campaigns_Show_Fail() {
	token := as.CreateUser("test@test.com", "test")

	req := as.JSON("/v1/campaigns/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Show_Other_User() {
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	// create campaign under token1
	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token1
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)

	// request campaign with token2
	req = as.JSON("/v1/campaigns/show/" + campaign.ID.String())
	req.Headers["Authorization"] = token2
	res = req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Destroy() {
	token := as.CreateUser("test@test.com", "test")

	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)

	// delete campaign
	req = as.JSON("/v1/campaigns/" + campaign.ID.String())
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Destroy_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// delete missing campaign id
	req := as.JSON("/v1/campaigns/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)

	// delete bad campaign id
	req = as.JSON("/v1/campaigns/abcde")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Destroy_Other_User() {
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token1
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.NotNil(campaign)

	// delete campaign
	req = as.JSON("/v1/campaigns/" + campaign.ID.String())
	req.Headers["Authorization"] = token2
	res = req.Delete()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Campaigns_List() {
	token := as.CreateUser("test@test.com", "test")

	campaignRequest1 := &CreateCampaignRequest{
		Name:        "Test Campaign 1",
		Description: "This is a test campaign.",
	}

	campaignRequest2 := &CreateCampaignRequest{
		Name:        "Test Campaign 2",
		Description: "This is another test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest1)
	as.Equal(http.StatusCreated, res.Code)
	res = req.Post(campaignRequest2)
	as.Equal(http.StatusCreated, res.Code)

	clr := &CampaignsListResponse{}
	req = as.JSON("/v1/campaigns/")
	req.Headers["Authorization"] = token
	res = req.Get()
	res.Bind(clr)

	as.Equal(2, len(clr.Campaigns))
}

func (as *ActionSuite) Test_Campaigns_Update() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	// create character data
	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.Equal("Test Campaign", campaign.Name)

	campaign.Name = "Fake Campaign"

	req = as.JSON("/v1/campaigns/" + campaign.ID.String())
	req.Headers["Authorization"] = token
	res = req.Put(campaign)
	as.Equal(http.StatusOK, res.Code)

	testCampaign := &models.Campaign{}
	res.Bind(testCampaign)
	as.Equal("Fake Campaign", testCampaign.Name)
}

func (as *ActionSuite) Test_Campaigns_Update_Fail() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	// create character data
	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.Equal("Test Campaign", campaign.Name)

	campaign.Name = "Fake Campaign"

	// fail on unknown uuid
	req = as.JSON("/v1/campaigns/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(campaign)
	as.Equal(http.StatusNotFound, res.Code)

	// fail on bad uuid
	req = as.JSON("/v1/campaigns/aaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(campaign)
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Campaigns_Update_Other_User() {
	// create valid user
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	// create character data
	campaignRequest := &CreateCampaignRequest{
		Name:        "Test Campaign",
		Description: "This is a test campaign.",
	}

	req := as.JSON("/v1/campaigns/new")
	req.Headers["Authorization"] = token1
	res := req.Post(campaignRequest)
	as.Equal(http.StatusCreated, res.Code)

	campaign := &models.Campaign{}
	res.Bind(campaign)
	as.Equal("Test Campaign", campaign.Name)

	campaign.Name = "Fake Campaign"

	req = as.JSON("/v1/campaigns/" + campaign.ID.String())
	req.Headers["Authorization"] = token2
	res = req.Put(campaign)
	as.Equal(http.StatusNotFound, res.Code)
}
