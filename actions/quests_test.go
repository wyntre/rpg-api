package actions

import (
	"github.com/gofrs/uuid"
	"github.com/wyntre/rpg_api/models"
	"net/http"
)

type CreateQuestRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CampaignID  uuid.UUID `json:"campaign_id"`
}

type QuestsListResponse struct {
	Quests models.Quests `json:quests`
}

func (as *ActionSuite) CreateQuest(name string, description string, campaign_id uuid.UUID, token string) *models.Quest {
	// create quest
	questRequest := &CreateQuestRequest{
		Name:        "Test Quest",
		Description: "This is a test quest.",
		CampaignID:  campaign_id,
	}

	req := as.JSON("/v1/quests/new")
	req.Headers["Authorization"] = token
	res := req.Post(questRequest)
	as.Equal(http.StatusCreated, res.Code)

	quest := &models.Quest{}
	res.Bind(quest)
	as.NotNil(quest)

	return quest
}

func (as *ActionSuite) Test_Quests_Create() {
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)

	// create quest
	questRequest := &CreateQuestRequest{
		Name:        "Test Quest",
		Description: "This is a test quest.",
		CampaignID:  campaign.ID,
	}

	req := as.JSON("/v1/quests/new")
	req.Headers["Authorization"] = token
	res := req.Post(questRequest)
	as.Equal(http.StatusCreated, res.Code)

	quest := &models.Quest{}
	res.Bind(quest)
	as.NotNil(quest)
}

func (as *ActionSuite) Test_Quests_Create_Fail() {
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)

	// without description
	questRequest := &CreateQuestRequest{
		Name:       "Test Quest",
		CampaignID: campaign.ID,
	}

	req := as.JSON("/v1/quests/new")
	req.Headers["Authorization"] = token
	res := req.Post(questRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)

	// without name
	questRequest = &CreateQuestRequest{
		Description: "This is a test quest.",
		CampaignID:  campaign.ID,
	}

	req = as.JSON("/v1/quests/new")
	req.Headers["Authorization"] = token
	res = req.Post(questRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)

	// without campaign_id
	questRequest = &CreateQuestRequest{
		Name:        "Test Quest",
		Description: "This is a test quest.",
	}

	req = as.JSON("/v1/quests/new")
	req.Headers["Authorization"] = token
	res = req.Post(questRequest)
	as.Equal(http.StatusUnprocessableEntity, res.Code)
}

func (as *ActionSuite) Test_Quests_Create_No_Token() {
	questRequest := &CreateQuestRequest{
		Name:        "Test Quest",
		Description: "This is a test quest.",
	}

	req := as.JSON("/v1/quests/new")
	res := req.Post(questRequest)
	as.Equal(http.StatusUnauthorized, res.Code)
}

func (as *ActionSuite) Test_Quests_Show() {
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)

	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token)

	req := as.JSON("/v1/quests/show/" + quest.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusOK, res.Code)

	test_quest := &models.Quest{}
	res.Bind(test_quest)
	as.Equal(quest.ID, test_quest.ID)
	as.Equal(quest.Name, test_quest.Name)
}

func (as *ActionSuite) Test_Quests_Show_Fail() {
	token := as.CreateUser("test@test.com", "test")

	req := as.JSON("/v1/quests/show/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Quests_Show_Other_User() {
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token1)

	// request quest with token2
	req := as.JSON("/v1/quests/" + quest.ID.String())
	req.Headers["Authorization"] = token2
	res := req.Get()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Quests_Destroy() {
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token)

	// delete quest
	req := as.JSON("/v1/quests/" + quest.ID.String())
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusOK, res.Code)
}

func (as *ActionSuite) Test_Quests_Destroy_Fail() {
	token := as.CreateUser("test@test.com", "test")

	// delete missing quest id
	req := as.JSON("/v1/quests/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)

	// delete bad quest id
	req = as.JSON("/v1/quests/abcde")
	req.Headers["Authorization"] = token
	res = req.Delete()
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Quests_Destroy_Other_User() {
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token1)

	// delete quest
	req := as.JSON("/v1/quests/" + quest.ID.String())
	req.Headers["Authorization"] = token2
	res := req.Delete()
	as.Equal(http.StatusNotFound, res.Code)
}

func (as *ActionSuite) Test_Quests_List() {
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
	as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token)
	as.CreateQuest("Test Quest 2", "Test Quest Description", campaign.ID, token)

	clr := &QuestsListResponse{}
	req := as.JSON("/v1/quests/" + campaign.ID.String())
	req.Headers["Authorization"] = token
	res := req.Get()
	res.Bind(clr)

	as.Equal(2, len(clr.Quests))
}

func (as *ActionSuite) Test_Quests_Update() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token)

	quest.Name = "Fake Quest"

	req := as.JSON("/v1/quests/" + quest.ID.String())
	req.Headers["Authorization"] = token
	res := req.Put(quest)
	as.Equal(http.StatusOK, res.Code)

	test_quest := &models.Quest{}
	res.Bind(test_quest)
	as.Equal("Fake Quest", test_quest.Name)
}

func (as *ActionSuite) Test_Quests_Update_Fail() {
	// create valid user
	token := as.CreateUser("test@test.com", "test")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token)

	quest.Name = "Fake Quest"

	// fail on unknown uuid
	req := as.JSON("/v1/quests/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	req.Headers["Authorization"] = token
	res := req.Put(quest)
	as.Equal(http.StatusNotFound, res.Code)

	// fail on bad uuid
	req = as.JSON("/v1/quests/aaaaaaaa")
	req.Headers["Authorization"] = token
	res = req.Put(quest)
	as.Equal(http.StatusInternalServerError, res.Code)
}

func (as *ActionSuite) Test_Quests_Update_Other_User() {
	// create valid user
	token1 := as.CreateUser("test@test.com", "test")
	token2 := as.CreateUser("fake@fake.com", "fake")

	campaign := as.CreateCampaign("Test Campagin", "Test Campaign Description", token1)
	quest := as.CreateQuest("Test Quest", "Test Quest Description", campaign.ID, token1)

	quest.Name = "Fake Quest"

	req := as.JSON("/v1/quests/" + quest.ID.String())
	req.Headers["Authorization"] = token2
	res := req.Put(quest)
	as.Equal(http.StatusNotFound, res.Code)
}
